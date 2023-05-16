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

func getPhase3TwoPart(item *pb.Phase3TwoPart) (*models.Phase3TwoPart, error) {
	if item == nil {
		return nil, nil
	}
	phase3TwoPart := &models.Phase3TwoPart{}
	err := util.FromMessage(item, phase3TwoPart)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return phase3TwoPart, nil
}

func getPhase3TwoParts(items []*pb.Phase3TwoPart) ([]*models.Phase3TwoPart, error) {
	phase3TwoParts := []*models.Phase3TwoPart{}
	for _, item := range items {
		phase3TwoPart, err := getPhase3TwoPart(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getPhase3TwoPart")
		}
		phase3TwoParts = append(phase3TwoParts, phase3TwoPart)
	}
	return phase3TwoParts, nil
}

func makePhase3TwoPart(item *models.Phase3TwoPart) (*pb.Phase3TwoPart, error) {
	if item == nil {
		return nil, nil
	}
	phase3TwoPart := &pb.Phase3TwoPart{}
	err := util.ToMessage(item, phase3TwoPart)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return phase3TwoPart, nil
}

func makePhase3TwoParts(items []models.Phase3TwoPart) ([]*pb.Phase3TwoPart, error) {
	phase3TwoParts := []*pb.Phase3TwoPart{}
	for _, item := range items {
		phase3TwoPart, err := makePhase3TwoPart(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makePhase3TwoPart")
		}
		phase3TwoParts = append(phase3TwoParts, phase3TwoPart)
	}
	return phase3TwoParts, nil
}

func (inst *RecommendService) Phase3TwoPartCreate(ctx context.Context, req *pb.Phase3TwoPartRequest) (*pb.MessageResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	phase3TwoPart, err := getPhase3TwoPart(req.Data)
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
	if combinedData.TenantId1 != req.Ctx.DomainId {
		return nil, status.Errorf(codes.PermissionDenied, "Tenant is not allow create phase3 two part")
	}
	phase3TwoPart.DomainId = combinedData.CombinedDataId
	phase3TwoPartTmp, err := inst.componentsContainer.Phase3TwoPartRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"PositionUser": phase3TwoPart.PositionUser,
		"TenantId":     phase3TwoPart.TenantId,
		"PositionItem": phase3TwoPart.PositionItem,
		"Part":         phase3TwoPart.Part,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if phase3TwoPartTmp == nil {
		phase3TwoPart.TenantId = req.Ctx.DomainId
		_, err = inst.componentsContainer.Phase3TwoPartRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), phase3TwoPart)
		if err != nil {
			return nil, errutil.Wrap(err, "Phase3TwoPartRepository.Create")
		}
	} else {
		phase3TwoPartTmp.ProcessedDataC1 = phase3TwoPart.ProcessedDataC1
		phase3TwoPartTmp.ProcessedDataC2 = phase3TwoPart.ProcessedDataC2
		_, err = inst.componentsContainer.Phase3TwoPartRepository().Update(userContext.SetDomainId(combinedData.CombinedDataId), phase3TwoPart, &entity.UpdateRequest{})
		if err != nil {
			return nil, errutil.Wrap(err, "Phase3TwoPartRepository.Update")
		}
	}
	return &pb.MessageResponse{}, nil
}

func (inst *RecommendService) Phase3TwoPartGet(ctx context.Context, req *pb.Phase3TwoPartRequest) (*pb.Phase3TwoPartResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	phase3TwoPart, err := getPhase3TwoPart(req.Data)
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
	if combinedData.TenantId2 != req.Ctx.DomainId {
		return nil, status.Errorf(codes.PermissionDenied, "Tenant is not allow get phase3 two part")
	}
	phase3TwoPartTmp, err := inst.componentsContainer.Phase3TwoPartRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
		"PositionUser": phase3TwoPart.PositionUser,
		"TenantId":     combinedData.TenantId1,
		"PositionItem": phase3TwoPart.PositionItem,
		"Part":         phase3TwoPart.Part,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "Phase3TwoPartRepository.FindByOneByAttribute")
	}
	item, err := makePhase3TwoPart(phase3TwoPartTmp)
	if err != nil {
		return nil, errutil.Wrap(err, "makeCombinedData")
	}
	return &pb.Phase3TwoPartResponse{
		Data: item,
	}, nil
}
