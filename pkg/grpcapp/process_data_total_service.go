package grpcapp

import (
	"context"
	"fmt"

	"github.com/thteam47/go-recommend-api/pkg/models"

	pb "github.com/thteam47/common/api/recommend-api"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (inst *RecommendService) ProcessDataTotal(ctx context.Context, req *pb.StringRequest) (*pb.MessageResponse, error) {
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
	totalNumberItem := combinedData.NumberItem1 + combinedData.NumberItem2

	for i := 1; i <= int(totalNumberItem); i++ {
		k := i
		tenantId := combinedData.TenantId1
		if i > int(combinedData.NumberItem1) {
			k = i - int(combinedData.NumberItem1)
			tenantId = combinedData.TenantId2
		}
		processDataSurvey, err := inst.componentsContainer.ProcessDataSurveyRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"PositionItem": k,
			"TenantId":     tenantId,
			"Part":         1,
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.FindByOneByAttribute")
		}
		if processDataSurvey == nil {
			return nil, status.Errorf(codes.NotFound, "process Data Survey not found")
		}

		err = inst.componentsContainer.ProcessDataTotalRepository().CreateAndUpdate(userContext.SetDomainId(combinedData.CombinedDataId), &models.ProcessDataTotal{
			DomainId:             combinedData.CombinedDataId,
			PositionItem:         int32(i),
			ProcessedData:        processDataSurvey.ProcessedData,
			PositionItemOriginal: int32(i),
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.CreateAndUpdate")
		}
	}

	for i := 1; i <= int(totalNumberItem); i++ {
		k := int(combinedData.NumberItem1) + i
		tenantId := combinedData.TenantId1
		if i > int(combinedData.NumberItem1) {
			k = int(combinedData.NumberItem2) + i - int(combinedData.NumberItem1)
			tenantId = combinedData.TenantId2
		}
		processDataSurvey, err := inst.componentsContainer.ProcessDataSurveyRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"PositionItem": k,
			"TenantId":     tenantId,
			"Part":         1,
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.FindByOneByAttribute")
		}
		if processDataSurvey == nil {
			return nil, status.Errorf(codes.NotFound, "process Data Survey not found")
		}
		err = inst.componentsContainer.ProcessDataTotalRepository().CreateAndUpdate(userContext.SetDomainId(combinedData.CombinedDataId), &models.ProcessDataTotal{
			DomainId:             combinedData.CombinedDataId,
			PositionItem:         totalNumberItem + int32(i),
			ProcessedData:        processDataSurvey.ProcessedData,
			PositionItemOriginal: int32(i),
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.CreateAndUpdate")
		}
	}

	for i := 1; i <= int(totalNumberItem); i++ {
		k := 2*int(combinedData.NumberItem1) + i
		tenantId := combinedData.TenantId1
		if i > int(combinedData.NumberItem1) {
			k = 2*int(combinedData.NumberItem2) + i - int(combinedData.NumberItem1)
			tenantId = combinedData.TenantId2
		}
		processDataSurvey, err := inst.componentsContainer.ProcessDataSurveyRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"PositionItem": k,
			"TenantId":     tenantId,
			"Part":         1,
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.FindByOneByAttribute")
		}
		if processDataSurvey == nil {
			return nil, status.Errorf(codes.NotFound, "process Data Survey not found")
		}
		err = inst.componentsContainer.ProcessDataTotalRepository().CreateAndUpdate(userContext.SetDomainId(combinedData.CombinedDataId), &models.ProcessDataTotal{
			DomainId:             combinedData.CombinedDataId,
			PositionItem:         2*totalNumberItem + int32(i),
			ProcessedData:        processDataSurvey.ProcessedData,
			PositionItemOriginal: int32(i),
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.CreateAndUpdate")
		}
	}

	count := 1
	countTenant1 := 1
	countTenant2 := 1
	countTenantTwoPart := 1
	for i := 1; i <= int(totalNumberItem)-1; i++ {
		for l := i + 1; l <= int(totalNumberItem); l++ {
			tenantId := combinedData.TenantId1
			if i > int(combinedData.NumberItem1) {
				tenantId = combinedData.TenantId2
			}
			processDataSurvey := &models.ProcessDataSurvey{}
			if (l <= int(combinedData.NumberItem1) && i <= int(combinedData.NumberItem1)) || (l > int(combinedData.NumberItem1) && i > int(combinedData.NumberItem1)) {
				countTenant := 3*int(combinedData.NumberItem1) + countTenant1
				if tenantId == combinedData.TenantId2 {
					countTenant = 3*int(combinedData.NumberItem2) + countTenant2
				}
				processDataSurvey, err = inst.componentsContainer.ProcessDataSurveyRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
					"PositionItem": countTenant,
					"TenantId":     tenantId,
					"Part":         1,
				})
				if err != nil {
					return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.FindByOneByAttribute")
				}
				if tenantId == combinedData.TenantId1 {
					countTenant1++
				} else {
					countTenant2++
				}
			} else {
				processDataSurvey, err = inst.componentsContainer.ProcessDataSurveyRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
					"PositionItem": countTenantTwoPart,
					"Part":         2,
				})
				if err != nil {
					return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.FindByOneByAttribute")
				}
				countTenantTwoPart++
			}

			if processDataSurvey == nil {
				return nil, status.Errorf(codes.NotFound, "process Data Survey not found")
			}
			err = inst.componentsContainer.ProcessDataTotalRepository().CreateAndUpdate(userContext.SetDomainId(combinedData.CombinedDataId), &models.ProcessDataTotal{
				DomainId:              combinedData.CombinedDataId,
				PositionItem:          3*totalNumberItem + int32(count),
				ProcessedData:         processDataSurvey.ProcessedData,
				PositionItemOriginal1: int32(i),
				PositionItemOriginal2: int32(l),
			})
			if err != nil {
				return nil, errutil.Wrap(err, "ProcessDataSurveyRepository.CreateAndUpdate")
			}

			count++
		}
	}

	fmt.Println(count)

	return &pb.MessageResponse{
		Ok: true,
	}, nil
}
