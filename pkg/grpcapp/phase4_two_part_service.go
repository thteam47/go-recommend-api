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

func getPhase4TwoPart(item *pb.Phase4TwoPart) (*models.Phase4TwoPart, error) {
	if item == nil {
		return nil, nil
	}
	phase4TwoPart := &models.Phase4TwoPart{}
	err := util.FromMessage(item, phase4TwoPart)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return phase4TwoPart, nil
}

func getPhase4TwoParts(items []*pb.Phase4TwoPart) ([]*models.Phase4TwoPart, error) {
	phase4TwoParts := []*models.Phase4TwoPart{}
	for _, item := range items {
		phase4TwoPart, err := getPhase4TwoPart(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getPhase4TwoPart")
		}
		phase4TwoParts = append(phase4TwoParts, phase4TwoPart)
	}
	return phase4TwoParts, nil
}

func makePhase4TwoPart(item *models.Phase4TwoPart) (*pb.Phase4TwoPart, error) {
	if item == nil {
		return nil, nil
	}
	phase4TwoPart := &pb.Phase4TwoPart{}
	err := util.ToMessage(item, phase4TwoPart)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return phase4TwoPart, nil
}

func makePhase4TwoParts(items []models.Phase4TwoPart) ([]*pb.Phase4TwoPart, error) {
	phase4TwoParts := []*pb.Phase4TwoPart{}
	for _, item := range items {
		phase4TwoPart, err := makePhase4TwoPart(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makePhase4TwoPart")
		}
		phase4TwoParts = append(phase4TwoParts, phase4TwoPart)
	}
	return phase4TwoParts, nil
}

func (inst *RecommendService) Phase4TwoPartCreate(ctx context.Context, req *pb.Phase4TwoPartRequest) (*pb.MessageResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	phase4TwoPart, err := getPhase4TwoPart(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getPhase4TwoPart")
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	if combinedData.TenantId2 != req.Ctx.DomainId {
		return nil, status.Errorf(codes.PermissionDenied, "Tenant is not allow create phase4 two part")
	}
	phase4TwoPart.DomainId = combinedData.CombinedDataId
	phase4TwoPartTmp, err := inst.componentsContainer.Phase4TwoPartRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"PositionUser": phase4TwoPart.PositionUser,
		"TenantId":     phase4TwoPart.TenantId,
		"PositionItem": phase4TwoPart.PositionItem,
		"Part":         phase4TwoPart.Part,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if phase4TwoPartTmp == nil {
		phase4TwoPart.TenantId = req.Ctx.DomainId
		_, err = inst.componentsContainer.Phase4TwoPartRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), phase4TwoPart)
		if err != nil {
			return nil, errutil.Wrap(err, "Phase4TwoPartRepository.Create")
		}
	} else {
		phase4TwoPartTmp.ProcessedDataR1 = phase4TwoPart.ProcessedDataR1
		phase4TwoPartTmp.ProcessedDataR2 = phase4TwoPart.ProcessedDataR2
		phase4TwoPartTmp.ProcessedDataR3 = phase4TwoPart.ProcessedDataR3
		_, err = inst.componentsContainer.Phase4TwoPartRepository().Update(userContext.SetDomainId(combinedData.CombinedDataId), phase4TwoPart, &entity.UpdateRequest{})
		if err != nil {
			return nil, errutil.Wrap(err, "Phase4TwoPartRepository.Update")
		}
	}
	return &pb.MessageResponse{}, nil
}

func (inst *RecommendService) Phase4TwoPartGet(ctx context.Context, req *pb.Phase4TwoPartRequest) (*pb.Phase4TwoPartResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	phase4TwoPart, err := getPhase4TwoPart(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getPhase4TwoPart")
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	if combinedData.TenantId1 != req.Ctx.DomainId {
		return nil, status.Errorf(codes.PermissionDenied, "Tenant is not allow get phase4 two part")
	}
	phase4TwoPartTmp, err := inst.componentsContainer.Phase4TwoPartRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"PositionUser": phase4TwoPart.PositionUser,
		"TenantId":     combinedData.TenantId2,
		"PositionItem": phase4TwoPart.PositionItem,
		"Part":         phase4TwoPart.Part,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "Phase4TwoPartRepository.FindByOneByAttribute")
	}
	item, err := makePhase4TwoPart(phase4TwoPartTmp)
	if err != nil {
		return nil, errutil.Wrap(err, "makeCombinedData")
	}
	return &pb.Phase4TwoPartResponse{
		Data: item,
	}, nil
}
