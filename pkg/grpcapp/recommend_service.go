package grpcapp

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strconv"

	"github.com/thteam47/common-libs/ellipticutil"
	"github.com/thteam47/common-libs/x509util"
	pb "github.com/thteam47/common/api/recommend-api"
	"github.com/thteam47/common/entity"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/component"
	"github.com/thteam47/go-recommend-api/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var curve = elliptic.P256()

type RecommendService struct {
	pb.RecommendServiceServer
	componentsContainer *component.ComponentsContainer
}

func NewRecommendService(componentsContainer *component.ComponentsContainer) *RecommendService {
	return &RecommendService{
		componentsContainer: componentsContainer,
	}
}
func getUser(item *pb.User) (*entity.User, error) {
	if item == nil {
		return nil, nil
	}
	user := &entity.User{}
	err := util.FromMessage(item, user)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return user, nil
}

func getUsers(items []*pb.User) ([]*entity.User, error) {
	users := []*entity.User{}
	for _, item := range items {
		user, err := getUser(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getUser")
		}
		users = append(users, user)
	}
	return users, nil
}

func makeUser(item *entity.User) (*pb.User, error) {
	if item == nil {
		return nil, nil
	}
	user := &pb.User{}
	err := util.ToMessage(item, user)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return user, nil
}

func makeUsers(items []entity.User) ([]*pb.User, error) {
	users := []*pb.User{}
	for _, item := range items {
		user, err := makeUser(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makeUser")
		}
		users = append(users, user)
	}
	return users, nil
}

func (inst *RecommendService) RecommendCbfGenarate(ctx context.Context, req *pb.RecommendRequest) (*pb.RecommendCbfResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}

	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	totalNumberItem := combinedData.NumberItem1 + combinedData.NumberItem2
	if req.Data == nil || len(req.Data) != int(totalNumberItem) {
		return nil, status.Errorf(codes.FailedPrecondition, "total item bad request")
	}
	cPri, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	cByte := new(big.Int).SetBytes(cPri.D.Bytes())

	keyInfo, err := inst.componentsContainer.KeyPublicUserRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"UserId":   req.UserId,
		"TenantId": req.Ctx.DomainId,
	})
	if keyInfo == nil {
		return nil, status.Errorf(codes.NotFound, "keyInfo not found")
	}
	publicKeyX, err := x509util.ExtractKeyPublic(keyInfo.KeyPublic)
	if err != nil {
		panic(err)
	}
	cjXx, cjXy := curve.ScalarMult(publicKeyX.X, publicKeyX.Y, cByte.Bytes())
	processDataResult := map[string]*pb.RecommendCbfResult12{}

	sumF3k := &ecdsa.PublicKey{
		Curve: curve,
		X:     nil,
		Y:     nil,
	}
	sumF4k := &ecdsa.PublicKey{
		Curve: curve,
		X:     nil,
		Y:     nil,
	}
	for k := 1; k <= int(totalNumberItem); k++ {
		sumF1k := &ecdsa.PublicKey{
			Curve: curve,
			X:     nil,
			Y:     nil,
		}
		sumF2k := &ecdsa.PublicKey{
			Curve: curve,
			X:     nil,
			Y:     nil,
		}
		for j := 1; j <= int(totalNumberItem); j++ {
			if j != k {
				processDataC1 := ""
				processDataC2 := ""
				if dataRecommendCbf, found := req.Data[strconv.Itoa(j)]; found {
					processDataC1 = dataRecommendCbf.ProcessDataC1
					processDataC2 = dataRecommendCbf.ProcessDataC2
				}
				pointC1, err := hex.DecodeString(processDataC1)
				if err != nil {
					panic(err)
				}
				C1X, C1Y := elliptic.Unmarshal(curve, pointC1)

				processDataSumilor, err := inst.componentsContainer.ProcessDataSumilorRepository().FindByOneByFindRequest(userContext.SetDomainId(combinedData.CombinedDataId), &entity.FindRequest{
					Filters: []entity.FindRequestFilter{
						entity.FindRequestFilter{
							Operator: entity.FindRequestFilterOperatorOr,
							SubFilters: []entity.FindRequestFilter{
								entity.FindRequestFilter{
									Operator: entity.FindRequestFilterOperatorAnd,
									SubFilters: []entity.FindRequestFilter{
										entity.FindRequestFilter{
											Key:      "DomainId",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    combinedData.CombinedDataId,
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal1",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(k),
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal2",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(j),
										},
									},
								},
								entity.FindRequestFilter{
									Operator: entity.FindRequestFilterOperatorAnd,
									SubFilters: []entity.FindRequestFilter{
										entity.FindRequestFilter{
											Key:      "DomainId",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    combinedData.CombinedDataId,
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal1",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(j),
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal2",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(k),
										},
									},
								},
							},
						},
					},
				})
				if err != nil {
					return nil, errutil.Wrap(err, "ProcessDataSumilorRepository.FindByOneByFindRequest")
				}
				if processDataSumilor == nil {
					return nil, errutil.Wrap(err, "processDataSumilor not found")
				}
				rij := new(big.Int).SetInt64(int64(processDataSumilor.ProcessedData))
				silCjx, silCjy := curve.ScalarMult(C1X, C1Y, rij.Bytes())
				sumF1k.X, sumF1k.Y = ellipticutil.AddPoint(curve, sumF1k.X, sumF1k.Y, silCjx, silCjy)

				//F2
				pointC2, err := hex.DecodeString(processDataC2)
				if err != nil {
					panic(err)
				}
				C2X, C2Y := elliptic.Unmarshal(curve, pointC2)
				sil2Cjx, sil2Cjy := curve.ScalarMult(C2X, C2Y, rij.Bytes())
				sumF2k.X, sumF2k.Y = ellipticutil.AddPoint(curve, sumF2k.X, sumF2k.Y, sil2Cjx, sil2Cjy)

				if k == 1 {
					//F3
					silGx, silGy := curve.ScalarBaseMult(rij.Bytes())
					sumF3k.X, sumF3k.Y = ellipticutil.AddPoint(curve, sumF3k.X, sumF3k.Y, silGx, silGy)
				}
			}
		}

		Fk1Decypt := elliptic.Marshal(curve, sumF1k.X, sumF1k.Y)
		Fk2Decypt := elliptic.Marshal(curve, sumF2k.X, sumF2k.Y)

		processDataResult[strconv.Itoa(k)] = &pb.RecommendCbfResult12{
			ProcessDataFk1: hex.EncodeToString(Fk1Decypt),
			ProcessDataFk2: hex.EncodeToString(Fk2Decypt),
		}
	}
	//F4
	sumF4k.X, sumF4k.Y = curve.ScalarBaseMult(cByte.Bytes())
	sumF3 := &ecdsa.PublicKey{
		Curve: curve,
		X:     nil,
		Y:     nil,
	}
	sumF3.X, sumF3.Y = ellipticutil.AddPoint(curve, sumF3k.X, sumF3k.Y, cjXx, cjXy)
	Fk3Decypt := elliptic.Marshal(curve, sumF3.X, sumF3.Y)
	Fk4Decypt := elliptic.Marshal(curve, sumF4k.X, sumF4k.Y)
	return &pb.RecommendCbfResponse{
		Data12: processDataResult,
		Data34: &pb.RecommendCbfResult34{
			ProcessDataFk3: hex.EncodeToString(Fk3Decypt),
			ProcessDataFk4: hex.EncodeToString(Fk4Decypt),
		},
	}, nil
}

func (inst *RecommendService) RecommendCfGenarate(ctx context.Context, req *pb.RecommendRequest) (*pb.RecommendCfResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}

	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	totalNumberItem := combinedData.NumberItem1 + combinedData.NumberItem2
	if req.Data == nil || len(req.Data) != int(totalNumberItem) {
		return nil, status.Errorf(codes.FailedPrecondition, "total item bad request")
	}
	cPri, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	cByte := new(big.Int).SetBytes(cPri.D.Bytes())

	keyInfo, err := inst.componentsContainer.KeyPublicUserRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"UserId":   req.UserId,
		"TenantId": req.Ctx.DomainId,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "KeyPublicUserRepository.FindByOneByAttribute")
	}
	if keyInfo == nil {
		return nil, status.Errorf(codes.NotFound, "keyInfo not found")
	}
	publicKeyX, err := x509util.ExtractKeyPublic(keyInfo.KeyPublic)
	if err != nil {
		return nil, errutil.Wrap(err, "x509util.ExtractKeyPublic")
	}
	cXx, cXy := curve.ScalarMult(publicKeyX.X, publicKeyX.Y, cByte.Bytes())
	processDataResult := map[string]*pb.RecommendCfResult910{}
	sumF11kP2 := &ecdsa.PublicKey{
		Curve: curve,
		X:     nil,
		Y:     nil,
	}
	for k := 1; k <= int(totalNumberItem); k++ {
		sumF9kP2 := &ecdsa.PublicKey{
			Curve: curve,
			X:     nil,
			Y:     nil,
		}
		sumF10kP2 := &ecdsa.PublicKey{
			Curve: curve,
			X:     nil,
			Y:     nil,
		}
		sumSumilor := 0
		processDataRtbk, err := inst.componentsContainer.ProcessDataRtbRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"DomainId":             combinedData.CombinedDataId,
			"PositionItemOriginal": int32(k),
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataRtbRepository.FindByOneByAttribute")
		}
		if processDataRtbk == nil {
			return nil, status.Errorf(codes.NotFound, "processDataRtb not found")
		}
		ck4Pri, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			panic(err)
		}
		ck4Byte := new(big.Int).SetBytes(ck4Pri.D.Bytes())

		ck5Pri, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			panic(err)
		}
		ck5Byte := new(big.Int).SetBytes(ck5Pri.D.Bytes())
		for j := 1; j <= int(totalNumberItem); j++ {
			if j != k {
				//F5
				processDataSumilor, err := inst.componentsContainer.ProcessDataSumilorRepository().FindByOneByFindRequest(userContext.SetDomainId(combinedData.CombinedDataId), &entity.FindRequest{
					Filters: []entity.FindRequestFilter{
						entity.FindRequestFilter{
							Operator: entity.FindRequestFilterOperatorOr,
							SubFilters: []entity.FindRequestFilter{
								entity.FindRequestFilter{
									Operator: entity.FindRequestFilterOperatorAnd,
									SubFilters: []entity.FindRequestFilter{
										entity.FindRequestFilter{
											Key:      "DomainId",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    combinedData.CombinedDataId,
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal1",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(k),
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal2",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(j),
										},
									},
								},
								entity.FindRequestFilter{
									Operator: entity.FindRequestFilterOperatorAnd,
									SubFilters: []entity.FindRequestFilter{
										entity.FindRequestFilter{
											Key:      "DomainId",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    combinedData.CombinedDataId,
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal1",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(j),
										},
										entity.FindRequestFilter{
											Key:      "PositionItemOriginal2",
											Operator: entity.FindRequestFilterOperatorEqualTo,
											Value:    int32(k),
										},
									},
								},
							},
						},
					},
				})
				if err != nil {
					return nil, errutil.Wrap(err, "ProcessDataSumilorRepository.FindByOneByFindRequest")
				}
				if processDataSumilor == nil {
					return nil, errutil.Wrap(err, "processDataSumilor not found")
				}

				sumSumilor += (int(processDataSumilor.ProcessedData))

				//F7
				processDataRtbj, err := inst.componentsContainer.ProcessDataRtbRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
					"DomainId":             combinedData.CombinedDataId,
					"PositionItemOriginal": int32(j),
				})
				if err != nil {
					return nil, errutil.Wrap(err, "ProcessDataRtbRepository.FindByOneByAttribute")
				}
				if processDataRtbj == nil {
					return nil, status.Errorf(codes.NotFound, "processDataRtbj not found")
				}

				RjSil := new(big.Int).SetInt64(int64(processDataRtbj.ProcessedData / 10))
				RjSilGx, RjSilGy := curve.ScalarBaseMult(RjSil.Bytes())

				ck5Xx, ck5Xy := curve.ScalarMult(publicKeyX.X, publicKeyX.Y, ck5Byte.Bytes())
				F7kx, F7ky := ellipticutil.AddPoint(curve, RjSilGx, RjSilGy, ck5Xx, ck5Xy)
				F8kx, F8ky := curve.ScalarBaseMult(ck5Byte.Bytes())

				//F9

				processDataC1 := ""
				processDataC2 := ""
				if dataRecommend, found := req.Data[strconv.Itoa(j)]; found {
					processDataC1 = dataRecommend.ProcessDataC1
					processDataC2 = dataRecommend.ProcessDataC2
				}
				pointC3, err := hex.DecodeString(processDataC1)
				if err != nil {
					panic(err)
				}
				C3jX, C3jY := elliptic.Unmarshal(curve, pointC3)
				// invert one point
				F7kyY := new(big.Int).Neg(F7ky)
				// point normalization
				F7kyYSub := new(big.Int).Mod(F7kyY, curve.Params().P)
				Ck3x, Ck3y := ellipticutil.AddPoint(curve, C3jX, C3jY, F7kx, F7kyYSub)

				silkj := new(big.Int).SetInt64(int64(processDataSumilor.ProcessedData))
				F9kP2x, F9kP2y := curve.ScalarMult(Ck3x, Ck3y, silkj.Bytes())
				sumF9kP2.X, sumF9kP2.Y = ellipticutil.AddPoint(curve, sumF9kP2.X, sumF9kP2.Y, F9kP2x, F9kP2y)

				//F10
				pointC4, err := hex.DecodeString(processDataC2)
				if err != nil {
					panic(err)
				}
				C4jX, C4jY := elliptic.Unmarshal(curve, pointC4)
				// invert one point
				F8kyY := new(big.Int).Neg(F8ky)
				// point normalization
				F8kyYSub := new(big.Int).Mod(F8kyY, curve.Params().P)
				Ck4F8x, Ck4F8y := ellipticutil.AddPoint(curve, C4jX, C4jY, F8kx, F8kyYSub)

				F10kP2x, F10kP2y := curve.ScalarMult(Ck4F8x, Ck4F8y, silkj.Bytes())
				sumF10kP2.X, sumF10kP2.Y = ellipticutil.AddPoint(curve, sumF10kP2.X, sumF10kP2.Y, F10kP2x, F10kP2y)
			}
		}

		//F5
		valueRkSil := int(processDataRtbk.ProcessedData /10) * sumSumilor
		RkSil := new(big.Int).SetInt64(int64(valueRkSil))
		RkSilGx, RkSilGy := curve.ScalarBaseMult(RkSil.Bytes())
		ck4Bytex, ck4Bytey := curve.ScalarMult(publicKeyX.X, publicKeyX.Y, ck4Byte.Bytes())
		F5kx, F5ky := ellipticutil.AddPoint(curve, RkSilGx, RkSilGy, ck4Bytex, ck4Bytey)
		F6kx, F6ky := curve.ScalarBaseMult(ck4Byte.Bytes())

		// F9
		F9kx, F9ky := ellipticutil.AddPoint(curve, F5kx, F5ky, sumF9kP2.X, sumF9kP2.Y)
		//F10
		F10kx, F10ky := ellipticutil.AddPoint(curve, F6kx, F6ky, sumF10kP2.X, sumF10kP2.Y)

		//gen
		Fk9Decypt := elliptic.Marshal(curve, F9kx, F9ky)
		Fk10Decypt := elliptic.Marshal(curve, F10kx, F10ky)

		//F11
		smlSumByte := new(big.Int).SetInt64(int64(sumSumilor))
		smlGx, smlGy := curve.ScalarBaseMult(smlSumByte.Bytes())

		sumF11kP2.X, sumF11kP2.Y = ellipticutil.AddPoint(curve, smlGx, smlGy, sumF11kP2.X, sumF11kP2.Y)

		processDataResult[strconv.Itoa(k)] = &pb.RecommendCfResult910{
			ProcessDataFk9:  hex.EncodeToString(Fk9Decypt),
			ProcessDataFk10: hex.EncodeToString(Fk10Decypt),
		}
	}
	F12x, F12y := curve.ScalarBaseMult(cByte.Bytes())
	F12Decypt := elliptic.Marshal(curve, F12x, F12y)

	sumF11kP2.X, sumF11kP2.Y = ellipticutil.AddPoint(curve, sumF11kP2.X, sumF11kP2.Y, cXx, cXy)
	F11Decypt := elliptic.Marshal(curve, sumF11kP2.X, sumF11kP2.Y)
	return &pb.RecommendCfResponse{
		Data910: processDataResult,
		Data1112: &pb.RecommendCfResult1112{
			ProcessDataFk12: hex.EncodeToString(F12Decypt),
			ProcessDataFk11: hex.EncodeToString(F11Decypt),
		},
	}, nil
}
