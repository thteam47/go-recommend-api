package grpcapp

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/thteam47/go-recommend-api/pkg/models"
	"github.com/thteam47/go-recommend-api/util"

	"github.com/thteam47/common-libs/ellipticutil"
	pb "github.com/thteam47/common/api/recommend-api"
	"github.com/thteam47/common/entity"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getProcessDataSurvey2(item *pb.ProcessDataSurvey2) (*models.ProcessDataSurvey2, error) {
	if item == nil {
		return nil, nil
	}
	processDataSurvey := &models.ProcessDataSurvey2{}
	err := util.FromMessage(item, processDataSurvey)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return processDataSurvey, nil
}

func (inst *RecommendService) ProcessDataSurvey(ctx context.Context, req *pb.StringRequest) (*pb.MessageResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Value)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	sPart := combinedData.SOnePart1
	if req.Value == combinedData.TenantId2 {
		sPart = combinedData.SOnePart2
	}
	if req.Ctx.Part == 2 {
		sPart = combinedData.STwoPart
	}
	max := 5 * int(combinedData.NumberUser) * 5
	timeStart := time.Now().UnixMilli()
	for j := 1; j <= int(sPart); j++ {
		sumAij := &ecdsa.PublicKey{
			Curve: curve,
			X:     nil,
			Y:     nil,
		}
		positionOriginal := int32(0)
		positionOriginal1 := int32(0)
		positionOriginal2 := int32(0)
		for i := 1; i <= int(combinedData.NumberUser); i++ {
			filter := map[string]interface{}{}
			if req.Ctx.Part == 1 {
				filter = map[string]interface{}{
					"PositionUser": i,
					"TenantId":     req.Value,
					"PositionItem": j,
					"Part":         req.Ctx.Part,
				}
			} else {
				filter = map[string]interface{}{
					"PositionUser": i,
					"PositionItem": j,
					"Part":         req.Ctx.Part,
				}
			}
			resultCard, err := inst.componentsContainer.ResultCardRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), filter)
			if err != nil {
				return nil, errutil.Wrap(err, "ResultCardRepository.FindByOneByAttribute")
			}
			if resultCard == nil {
				return nil, status.Errorf(codes.FailedPrecondition, "ResultCard data not found")
			}
			positionOriginal = resultCard.PositionItemOriginal
			positionOriginal1 = resultCard.PositionItemOriginal1
			positionOriginal2 = resultCard.PositionItemOriginal2
			point, err := hex.DecodeString(resultCard.ProcessedData)
			if err != nil {
				panic(err)
			}
			AijX, AijY := elliptic.Unmarshal(curve, point)
			sumAij.X, sumAij.Y = ellipticutil.AddPoint(curve, sumAij.X, sumAij.Y, AijX, AijY)
		}

		sumRating := 0

		for x := 1; x <= max; x++ {
			xByte := new(big.Int).SetInt64(int64(x))
			xX, xY := curve.ScalarBaseMult(xByte.Bytes())
			if xX != nil && xY != nil && sumAij.X != nil && sumAij.Y != nil {
				if xX.Cmp(sumAij.X) == 0 && xY.Cmp(sumAij.Y) == 0 {
					sumRating = x
				}
			}
		}
		processDataSurvey := &models.ProcessDataSurvey{
			PositionItem:          int32(j),
			ProcessedData:         int32(sumRating),
			DomainId:              combinedData.CombinedDataId,
			TenantId:              req.Value,
			Part:                  req.Ctx.Part,
			PositionItemOriginal:  positionOriginal,
			PositionItemOriginal1: positionOriginal1,
			PositionItemOriginal2: positionOriginal2,
		}

		filterProcess := map[string]interface{}{}
		if req.Ctx.Part == 1 {
			filterProcess = map[string]interface{}{
				"PositionItem": int32(j),
				"TenantId":     req.Value,
				"Part":         req.Ctx.Part,
			}
		} else {
			filterProcess = map[string]interface{}{
				"PositionItem": int32(j),
				"Part":         req.Ctx.Part,
			}
		}
		processDataSurveyTmp, err := inst.componentsContainer.ProcessDataSurveyRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), filterProcess)
		if err != nil {
			return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
		}
		if processDataSurveyTmp == nil {
			_, err = inst.componentsContainer.ProcessDataSurveyRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), processDataSurvey)
			if err != nil {
				return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.Create")
			}
		} else {
			processDataSurveyTmp.ProcessedData = int32(sumRating)
			processDataSurveyTmp.PositionItemOriginal = processDataSurvey.PositionItemOriginal
			processDataSurveyTmp.PositionItemOriginal1 = processDataSurvey.PositionItemOriginal1
			processDataSurveyTmp.PositionItemOriginal2 = processDataSurvey.PositionItemOriginal2
			if req.Ctx.Part == 2 {
				processDataSurvey2Tmp, err := inst.componentsContainer.ProcessDataSurvey2Repository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
					"PositionItem": int32(j),
					"Part":         2,
				})
				if err != nil {
					return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
				}
				if processDataSurvey2Tmp != nil {
					processDataSurveyTmp.ProcessedData = processDataSurvey2Tmp.ProcessedData
					processDataSurveyTmp.PositionItemOriginal1 = processDataSurvey2Tmp.PositionItemOriginal1
					processDataSurveyTmp.PositionItemOriginal2 = processDataSurvey2Tmp.PositionItemOriginal2
				}
			}
			_, err = inst.componentsContainer.ProcessDataSurveyRepository().Update(userContext.SetDomainId(combinedData.CombinedDataId), processDataSurveyTmp, &entity.UpdateRequest{})
			if err != nil {
				return nil, errutil.Wrap(err, "ResultCardRepository.Update")
			}
		}
	}

	timeEnd := time.Now().UnixMilli()
	dentaTime := timeEnd - timeStart
	fmt.Printf("Time logarit: %d \n", dentaTime)
	return &pb.MessageResponse{}, nil
}

func (inst *RecommendService) ProcessDataSurveyCreate2(ctx context.Context, req *pb.ProcessDataSurvey2Request) (*pb.MessageResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	processDataSurvey, err := getProcessDataSurvey2(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getPhase3TwoPart")
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	processDataSurvey.DomainId = combinedData.CombinedDataId
	processDataSurveyTmp, err := inst.componentsContainer.ProcessDataSurvey2Repository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"PositionItem": processDataSurvey.PositionItem,
		"Part":         processDataSurvey.Part,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if processDataSurveyTmp == nil {
		_, err = inst.componentsContainer.ProcessDataSurvey2Repository().Create(userContext.SetDomainId(combinedData.CombinedDataId), processDataSurvey)
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurvey2Repository.Create")
		}
	} else {
		processDataSurveyTmp.ProcessedData = processDataSurvey.ProcessedData
		processDataSurveyTmp.PositionItemOriginal = processDataSurvey.PositionItemOriginal
		processDataSurveyTmp.PositionItemOriginal1 = processDataSurvey.PositionItemOriginal1
		processDataSurveyTmp.PositionItemOriginal2 = processDataSurvey.PositionItemOriginal2
		_, err = inst.componentsContainer.ProcessDataSurvey2Repository().Update(userContext.SetDomainId(combinedData.CombinedDataId), processDataSurveyTmp, &entity.UpdateRequest{})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurvey2Repository.Update")
		}
	}
	return &pb.MessageResponse{}, nil
}
