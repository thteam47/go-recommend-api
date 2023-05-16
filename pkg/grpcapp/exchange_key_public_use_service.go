package grpcapp

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"strconv"

	"github.com/thteam47/common-libs/ecdsautil"
	"github.com/thteam47/common-libs/x509util"
	pb "github.com/thteam47/common/api/recommend-api"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
	"github.com/thteam47/go-recommend-api/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getKeyPublicUse(item *pb.KeyPublicUse) (*models.KeyPublicUse, error) {
	if item == nil {
		return nil, nil
	}
	keyPublicUser := &models.KeyPublicUse{}
	err := util.FromMessage(item, keyPublicUser)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return keyPublicUser, nil
}

func getKeyPublicUses(items []*pb.KeyPublicUse) ([]*models.KeyPublicUse, error) {
	keyPublicUsers := []*models.KeyPublicUse{}
	for _, item := range items {
		keyPublicUser, err := getKeyPublicUse(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getKeyPublicUse")
		}
		keyPublicUsers = append(keyPublicUsers, keyPublicUser)
	}
	return keyPublicUsers, nil
}

func makeKeyPublicUse(item *models.KeyPublicUse) (*pb.KeyPublicUse, error) {
	if item == nil {
		return nil, nil
	}
	keyPublicUser := &pb.KeyPublicUse{}
	err := util.ToMessage(item, keyPublicUser)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return keyPublicUser, nil
}

func makeKeyPublicUses(items []models.KeyPublicUse) ([]*pb.KeyPublicUse, error) {
	keyPublicUsers := []*pb.KeyPublicUse{}
	for _, item := range items {
		keyPublicUser, err := makeKeyPublicUse(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makeKeyPublicUse")
		}
		keyPublicUsers = append(keyPublicUsers, keyPublicUser)
	}
	return keyPublicUsers, nil
}

func (inst *RecommendService) KeyPublicUseGet(ctx context.Context, req *pb.StringRequest) (*pb.KeyPublicUseResponse, error) {
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
	tenantId := combinedData.TenantId1
	if req.Ctx.DomainId == combinedData.TenantId2 {
		tenantId = combinedData.TenantId2
	}
	nkOnePart := combinedData.NkOnePart1
	if req.Ctx.DomainId == combinedData.TenantId2 {
		nkOnePart = combinedData.NkOnePart2
	}
	positionItem, err := strconv.Atoi(req.Value)
	if err != nil {
		return nil, errutil.Wrap(err, "strconv.Atoi")
	}
	keyPublicUse := &models.KeyPublicUse{}
	if req.Ctx.Part == 1 {
		keyPublicUse, err = inst.componentsContainer.KeyPublicUseRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"Part":         req.Ctx.Part,
			"TenantId":     req.Ctx.DomainId,
			"PositionItem": positionItem,
		})
		if err != nil {
			return nil, errutil.Wrap(err, "KeyPublicUseRepository.Update")
		}
		if keyPublicUse == nil {

			if positionItem <= 0 || positionItem > int(nkOnePart) {
				return nil, status.Errorf(codes.FailedPrecondition, "Position item not found")
			}
			publicKeyJ := &ecdsa.PublicKey{
				Curve: curve,
				X:     nil,
				Y:     nil,
			}
			for i := 1; i <= int(combinedData.NumberUser); i++ {
				keyPublicItemTmp, err := inst.componentsContainer.KeyPublicItemRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
					"PositionUser": i,
					"TenantId":     tenantId,
					"PositionItem": positionItem,
					"Part":         req.Ctx.Part,
				})
				if err != nil {
					return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
				}
				publicKeyIJ, err := x509util.ExtractKeyPublic(keyPublicItemTmp.KeyPublic)
				if err != nil {
					fmt.Println(err)
				}
				publicKeyJ = ecdsautil.AddPublicKeys(publicKeyJ, publicKeyIJ)
			}
			publicKeyGen, err := x509util.GenerateKeyPublic(publicKeyJ)
			if err != nil {
				panic(err)
			}
			keyPublicUse = &models.KeyPublicUse{
				Part:         req.Ctx.Part,
				TenantId:     req.Ctx.DomainId,
				PositionItem: int32(positionItem),
				KeyPublic:    publicKeyGen,
			}
			keyPublicUse, err = inst.componentsContainer.KeyPublicUseRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), keyPublicUse)
			if err != nil {
				return nil, errutil.Wrap(err, "KeyPublicUserRepository.Create")
			}
		}
	} else if req.Ctx.Part == 2 {
		keyPublicUse, err = inst.componentsContainer.KeyPublicUseRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"Part":         req.Ctx.Part,
			"PositionItem": positionItem,
		})
		if err != nil {
			return nil, errutil.Wrap(err, "KeyPublicUseRepository.Update")
		}
		if keyPublicUse == nil {
			if positionItem <= 0 || positionItem > int(combinedData.SkTwoPart) {
				return nil, status.Errorf(codes.FailedPrecondition, "Position item not found")
			}
			j := 1
			isStopj := false
			for t := 1; t <= int(combinedData.NkTwoPart); t++ {
				if isStopj {
					break
				}
				for k := 1; k <= int(combinedData.NkTwoPart); k++ {
					if j != positionItem {
						j++
						continue
					}
					publicKeyJ := &ecdsa.PublicKey{
						Curve: curve,
						X:     nil,
						Y:     nil,
					}
					for i := 1; i <= int(combinedData.NumberUser); i++ {
						keyPublicItemTmpIJU, err := inst.componentsContainer.KeyPublicItemRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
							"PositionUser": i,
							"PositionItem": t,
							"Part":         req.Ctx.Part,
							"TenantId":     combinedData.TenantId1,
						})
						if err != nil {
							return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
						}
						if keyPublicItemTmpIJU == nil {
							return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("keyPublicItemTmpIJU item not found: User %d, position: %d, part: %d", i, positionItem, req.Ctx.Part))
						}
						publicKeyIJU, err := x509util.ExtractKeyPublic(keyPublicItemTmpIJU.KeyPublic)
						if err != nil {
							fmt.Println(err)
						}
						publicKeyJ = ecdsautil.AddPublicKeys(publicKeyJ, publicKeyIJU)
						keyPublicItemTmpIJV, err := inst.componentsContainer.KeyPublicItemRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
							"PositionUser": i,
							"PositionItem": k,
							"Part":         req.Ctx.Part,
							"TenantId":     combinedData.TenantId2,
						})
						if err != nil {
							return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("keyPublicItemTmpIJU item not found: User %d, position: %d, part: %d", i, positionItem, req.Ctx.Part))
						}
						publicKeyIJV, err := x509util.ExtractKeyPublic(keyPublicItemTmpIJV.KeyPublic)
						if err != nil {
							fmt.Println(err)
						}
						publicKeyJ = ecdsautil.AddPublicKeys(publicKeyJ, publicKeyIJV)
					}
					publicKeyGen, err := x509util.GenerateKeyPublic(publicKeyJ)
					if err != nil {
						panic(err)
					}
					keyPublicUse = &models.KeyPublicUse{
						Part:         req.Ctx.Part,
						PositionItem: int32(positionItem),
						KeyPublic:    publicKeyGen,
					}
					keyPublicUse, err = inst.componentsContainer.KeyPublicUseRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), keyPublicUse)
					if err != nil {
						return nil, errutil.Wrap(err, "KeyPublicUserRepository.Create")
					}
					isStopj = true
					break
				}
			}
		}
	} else {
		return nil, status.Errorf(codes.FailedPrecondition, "Part not found")
	}
	item, err := makeKeyPublicUse(keyPublicUse)
	if err != nil {
		return nil, errutil.Wrap(err, "makeCombinedData")
	}
	return &pb.KeyPublicUseResponse{
		Data: item,
	}, nil
}
