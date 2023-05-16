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

func getResultCard(item *pb.ResultCard) (*models.ResultCard, error) {
	if item == nil {
		return nil, nil
	}
	resultCard := &models.ResultCard{}
	err := util.FromMessage(item, resultCard)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return resultCard, nil
}

func getResultCards(items []*pb.ResultCard) ([]*models.ResultCard, error) {
	resultCards := []*models.ResultCard{}
	for _, item := range items {
		resultCard, err := getResultCard(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getResultCard")
		}
		resultCards = append(resultCards, resultCard)
	}
	return resultCards, nil
}

func makeResultCard(item *models.ResultCard) (*pb.ResultCard, error) {
	if item == nil {
		return nil, nil
	}
	resultCard := &pb.ResultCard{}
	err := util.ToMessage(item, resultCard)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return resultCard, nil
}

func makeResultCards(items []models.ResultCard) ([]*pb.ResultCard, error) {
	resultCards := []*pb.ResultCard{}
	for _, item := range items {
		resultCard, err := makeResultCard(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makeResultCard")
		}
		resultCards = append(resultCards, resultCard)
	}
	return resultCards, nil
}

func (inst *RecommendService) ResultCardCreate(ctx context.Context, req *pb.ResultCardRequest) (*pb.MessageResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	resultCard, err := getResultCard(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getResultCard")
	}
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Ctx.DomainId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	filter := map[string]interface{}{}
	if resultCard.Part == 1 {
		filter = map[string]interface{}{
			"UserId":       resultCard.UserId,
			"TenantId":     resultCard.TenantId,
			"PositionItem": resultCard.PositionItem,
			"Part":         resultCard.Part,
		}
	} else {
		filter = map[string]interface{}{
			"UserId":       resultCard.UserId,
			"PositionItem": resultCard.PositionItem,
			"Part":         resultCard.Part,
		}
	}
	resultCard.DomainId = combinedData.CombinedDataId
	resultCardTmp, err := inst.componentsContainer.ResultCardRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), filter)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if resultCardTmp == nil {
		if resultCard.Part == 1 {
			resultCard.TenantId = req.Ctx.DomainId
		} else {
			resultCard.TenantId = ""
		}
		_, err = inst.componentsContainer.ResultCardRepository().Create(userContext.SetDomainId(combinedData.CombinedDataId), resultCard)
		if err != nil {
			return nil, errutil.Wrap(err, "ResultCardRepository.Create")
		}
	} else {
		resultCardTmp.ProcessedData = resultCard.ProcessedData
		resultCardTmp.PositionItemOriginal = resultCard.PositionItemOriginal
		resultCardTmp.PositionItemOriginal1 = resultCard.PositionItemOriginal1
		resultCardTmp.PositionItemOriginal2 = resultCard.PositionItemOriginal2
		_, err = inst.componentsContainer.ResultCardRepository().Update(userContext.SetDomainId(combinedData.CombinedDataId), resultCardTmp, &entity.UpdateRequest{})
		if err != nil {
			return nil, errutil.Wrap(err, "ResultCardRepository.Update")
		}
	}
	return &pb.MessageResponse{}, nil
}
