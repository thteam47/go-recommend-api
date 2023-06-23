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

func getKeyPublicUser(item *pb.KeyPublicUser) (*models.KeyPublicUser, error) {
	if item == nil {
		return nil, nil
	}
	keyPublicUser := &models.KeyPublicUser{}
	err := util.FromMessage(item, keyPublicUser)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return keyPublicUser, nil
}

func getKeyPublicUsers(items []*pb.KeyPublicUser) ([]*models.KeyPublicUser, error) {
	keyPublicUsers := []*models.KeyPublicUser{}
	for _, item := range items {
		keyPublicUser, err := getKeyPublicUser(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getKeyPublicUser")
		}
		keyPublicUsers = append(keyPublicUsers, keyPublicUser)
	}
	return keyPublicUsers, nil
}

func makeKeyPublicUser(item *models.KeyPublicUser) (*pb.KeyPublicUser, error) {
	if item == nil {
		return nil, nil
	}
	keyPublicUser := &pb.KeyPublicUser{}
	err := util.ToMessage(item, keyPublicUser)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return keyPublicUser, nil
}

func makeKeyPublicUsers(items []models.KeyPublicUser) ([]*pb.KeyPublicUser, error) {
	keyPublicUsers := []*pb.KeyPublicUser{}
	for _, item := range items {
		keyPublicUser, err := makeKeyPublicUser(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makeKeyPublicUser")
		}
		keyPublicUsers = append(keyPublicUsers, keyPublicUser)
	}
	return keyPublicUsers, nil
}

func (inst *RecommendService) KeyPublicUserReceive(ctx context.Context, req *pb.KeyPublicUserRequest) (*pb.MessageResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	keyPublicUser, err := getKeyPublicUser(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getKeyPublicUser")
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	keyPublicUser.DomainId = combinedData.CombinedDataId
	keyPublicUserTmp, err := inst.componentsContainer.KeyPublicUserRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"UserId":   keyPublicUser.UserId,
		"TenantId": req.Ctx.DomainId,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if keyPublicUserTmp == nil {
		keyPublicUser.TenantId = req.Ctx.DomainId
		_, err = inst.componentsContainer.KeyPublicUserRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), keyPublicUser)
		if err != nil {
			return nil, errutil.Wrap(err, "KeyPublicUserRepository.Create")
		}
	} else {
		keyPublicUserTmp.KeyPublic = keyPublicUser.KeyPublic
		_, err = inst.componentsContainer.KeyPublicUserRepository().Update(userContext.SetDomainId(combinedData.CombinedDataId), keyPublicUserTmp, &entity.UpdateRequest{})
		if err != nil {
			return nil, errutil.Wrap(err, "KeyPublicUserRepository.Update")
		}
	}
	return &pb.MessageResponse{}, nil
}

func (inst *RecommendService) KeyPublicUserGet(ctx context.Context, req *pb.KeyPublicUserRequest) (*pb.KeyPublicUserResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	keyPublicUser, err := getKeyPublicUser(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getKeyPublicUser")
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	keyPublicUserTmp, err := inst.componentsContainer.KeyPublicUserRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"PositionUser": keyPublicUser.PositionUser,
		"TenantId":     req.Ctx.DomainId,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	item, err := makeKeyPublicUser(keyPublicUserTmp)
	if err != nil {
		return nil, errutil.Wrap(err, "makeCombinedData")
	}
	return &pb.KeyPublicUserResponse{
		Data: item,
	}, nil
}
