package grpcapp

import (
	"context"

	pb "github.com/thteam47/common/api/recommend-api"
	"github.com/thteam47/common/entity"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
	"github.com/thteam47/go-recommend-api/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getKeyPublicItem(item *pb.KeyPublicItem) (*models.KeyPublicItem, error) {
	if item == nil {
		return nil, nil
	}
	keyPublicUser := &models.KeyPublicItem{}
	err := util.FromMessage(item, keyPublicUser)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return keyPublicUser, nil
}

func getKeyPublicItems(items []*pb.KeyPublicItem) ([]*models.KeyPublicItem, error) {
	keyPublicUsers := []*models.KeyPublicItem{}
	for _, item := range items {
		keyPublicUser, err := getKeyPublicItem(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getKeyPublicItem")
		}
		keyPublicUsers = append(keyPublicUsers, keyPublicUser)
	}
	return keyPublicUsers, nil
}

func makeKeyPublicItem(item *models.KeyPublicItem) (*pb.KeyPublicItem, error) {
	if item == nil {
		return nil, nil
	}
	keyPublicUser := &pb.KeyPublicItem{}
	err := util.ToMessage(item, keyPublicUser)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return keyPublicUser, nil
}

func makeKeyPublicItems(items []models.KeyPublicItem) ([]*pb.KeyPublicItem, error) {
	keyPublicUsers := []*pb.KeyPublicItem{}
	for _, item := range items {
		keyPublicUser, err := makeKeyPublicItem(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makeKeyPublicItem")
		}
		keyPublicUsers = append(keyPublicUsers, keyPublicUser)
	}
	return keyPublicUsers, nil
}

func (inst *RecommendService) KeyPublicItemReceive(ctx context.Context, req *pb.KeyPublicItemRequest) (*pb.MessageResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	keyPublicItem, err := getKeyPublicItem(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getKeyPublicItem")
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	keyPublicItem.DomainId = combinedData.CombinedDataId
	keyPublicItemTmp, err := inst.componentsContainer.KeyPublicItemRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"UserId":       keyPublicItem.UserId,
		"TenantId":     req.Ctx.DomainId,
		"PositionItem": keyPublicItem.PositionItem,
		"Part":         keyPublicItem.Part,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if keyPublicItemTmp == nil {
		keyPublicItem.TenantId = req.Ctx.DomainId
		_, err = inst.componentsContainer.KeyPublicItemRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), keyPublicItem)
		if err != nil {
			return nil, errutil.Wrap(err, "KeyPublicItemRepository.Create")
		}
	} else {
		keyPublicItemTmp.KeyPublic = keyPublicItem.KeyPublic
		_, err = inst.componentsContainer.KeyPublicItemRepository().Update(userContext.SetDomainId(combinedData.CombinedDataId), keyPublicItem, &entity.UpdateRequest{})
		if err != nil {
			return nil, errutil.Wrap(err, "KeyPublicItemRepository.Update")
		}
	}
	return &pb.MessageResponse{}, nil
}
