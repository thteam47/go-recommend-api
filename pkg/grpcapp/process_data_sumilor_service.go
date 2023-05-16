package grpcapp

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/thteam47/go-recommend-api/pkg/models"

	pb "github.com/thteam47/common/api/recommend-api"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (inst *RecommendService) ProcessDataSumilor(ctx context.Context, req *pb.StringRequest) (*pb.MessageResponse, error) {
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
	l := 1
	for j := 1; j <= int(totalNumberItem)-1; j++ {
		for k := j + 1; k <= int(totalNumberItem); k++ {
			processDataTotalTu, err := inst.componentsContainer.ProcessDataTotalRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
				"DomainId":     combinedData.CombinedDataId,
				"PositionItem": 3*totalNumberItem + int32(l),
			})
			if err != nil {
				return nil, errutil.Wrap(err, "ProcessDataTotalRepository.FindByOneByAttribute")
			}
			processDataTotalMau1, err := inst.componentsContainer.ProcessDataTotalRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
				"DomainId":     combinedData.CombinedDataId,
				"PositionItem": 2*totalNumberItem + int32(j),
			})
			if err != nil {
				return nil, errutil.Wrap(err, "ProcessDataTotalRepository.FindByOneByAttribute")
			}
			processDataTotalMau2, err := inst.componentsContainer.ProcessDataTotalRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
				"DomainId":     combinedData.CombinedDataId,
				"PositionItem": 2*totalNumberItem + int32(k),
			})
			if err != nil {
				return nil, errutil.Wrap(err, "ProcessDataTotalRepository.FindByOneByAttribute")
			}
			sumilor := float64(processDataTotalTu.ProcessedData) / (math.Sqrt(float64(processDataTotalMau1.ProcessedData)) * math.Sqrt(float64(processDataTotalMau2.ProcessedData)))
			numStr, err := strconv.ParseFloat(fmt.Sprintf("%.2f", sumilor), 64)
			if err != nil {
				return nil, errutil.Wrap(err, "strconv.ParseFloat")
			}
			err = inst.componentsContainer.ProcessDataSumilorRepository().CreateAndUpdate(userContext.SetDomainId(combinedData.CombinedDataId), &models.ProcessDataSumilor{
				DomainId:              combinedData.CombinedDataId,
				PositionItemOriginal1: int32(j),
				PositionItemOriginal2: int32(k),
				ProcessedData:         int32(numStr * 100),
			})
			if err != nil {
				return nil, errutil.Wrap(err, "ProcessDataSumilorRepository.CreateAndUpdate")
			}
		}
	}

	for j := 1; j <= int(totalNumberItem); j++ {
		processDataTotalTu, err := inst.componentsContainer.ProcessDataTotalRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"DomainId":     combinedData.CombinedDataId,
			"PositionItem": int32(j),
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataTotalRepository.FindByOneByAttribute")
		}

		processDataTotalMau, err := inst.componentsContainer.ProcessDataTotalRepository().FindByOneByAttribute(userContext.SetDomainId(combinedData.CombinedDataId), map[string]interface{}{
			"DomainId":     combinedData.CombinedDataId,
			"PositionItem": totalNumberItem + int32(j),
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataTotalRepository.FindByOneByAttribute")
		}
		Rtb := float64(processDataTotalTu.ProcessedData) / float64(processDataTotalMau.ProcessedData)
		numStr, err := strconv.ParseFloat(fmt.Sprintf("%.2f", Rtb), 64)
		if err != nil {
			return nil, errutil.Wrap(err, "strconv.ParseFloat")
		}

		err = inst.componentsContainer.ProcessDataRtbRepository().CreateAndUpdate(userContext.SetDomainId(combinedData.CombinedDataId), &models.ProcessDataRtb{
			DomainId:             combinedData.CombinedDataId,
			PositionItemOriginal: int32(j),
			ProcessedData:        int32(numStr * 100),
		})
		if err != nil {
			return nil, errutil.Wrap(err, "ProcessDataSumilorRepository.CreateAndUpdate")
		}
	}

	return &pb.MessageResponse{
		Ok: true,
	}, nil
}
