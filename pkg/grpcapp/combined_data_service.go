package grpcapp

import (
	"context"
	"math"

	"github.com/thteam47/common/pkg/entityutil"

	pb "github.com/thteam47/common/api/recommend-api"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
	"github.com/thteam47/go-recommend-api/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getCombinedData(item *pb.CombinedData) (*models.CombinedData, error) {
	if item == nil {
		return nil, nil
	}
	combinedData := &models.CombinedData{}
	err := util.FromMessage(item, combinedData)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return combinedData, nil
}

func getCombinedDatas(items []*pb.CombinedData) ([]*models.CombinedData, error) {
	combinedDatas := []*models.CombinedData{}
	for _, item := range items {
		combinedData, err := getCombinedData(item)
		if err != nil {
			return nil, errutil.Wrap(err, "getCombinedData")
		}
		combinedDatas = append(combinedDatas, combinedData)
	}
	return combinedDatas, nil
}

func makeCombinedData(item *models.CombinedData) (*pb.CombinedData, error) {
	if item == nil {
		return nil, nil
	}
	combinedData := &pb.CombinedData{}
	err := util.ToMessage(item, combinedData)
	if err != nil {
		return nil, errutil.Wrap(err, "ToMessage")
	}
	return combinedData, nil
}

func makeCombinedDatas(items []models.CombinedData) ([]*pb.CombinedData, error) {
	combinedDatas := []*pb.CombinedData{}
	for _, item := range items {
		combinedData, err := makeCombinedData(&item)
		if err != nil {
			return nil, errutil.Wrap(err, "makeCombinedData")
		}
		combinedDatas = append(combinedDatas, combinedData)
	}
	return combinedDatas, nil
}

func (inst *RecommendService) CombinedDataCreate(ctx context.Context, req *pb.CombinedDataRequest) (*pb.CombinedDataResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	combinedData, err := getCombinedData(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getCombinedData")
	}
	if !entityutil.ServiceOrAdminRole(userContext) {
		tenant, err := inst.componentsContainer.CustomerService().TenantFindById(userContext, combinedData.TenantId1)
		if err != nil {
			return nil, errutil.Wrap(err, "CustomerService.TenantFindById")
		}
		if tenant == nil {
			return nil, status.Errorf(codes.NotFound, "Tenant not found")
		}
		customerID, err := entityutil.GetUserId(userContext)
		if err != nil {
			return nil, errutil.Wrap(err, "entityutil.GetUserId")
		}
		if customerID != tenant.CustomerId {
			return nil, status.Errorf(codes.PermissionDenied, "PermissionDenied")
		}
		countUsersTenant1, err := inst.componentsContainer.IdentityService().GetCountUsers(userContext, combinedData.TenantId1)
		if err != nil {
			return nil, errutil.Wrap(err, "IdentityService.GetCountUsers")
		}
		countUsersTenant2, err := inst.componentsContainer.IdentityService().GetCountUsers(userContext, combinedData.TenantId2)
		if err != nil {
			return nil, errutil.Wrap(err, "IdentityService.GetCountUsers")
		}
		if combinedData.NumberUser > countUsersTenant1 || combinedData.NumberUser > countUsersTenant2 || combinedData.NumberUser < 0 {
			return nil, status.Errorf(codes.FailedPrecondition, "The number of users participating in the combination is not satisfied")
		}
	}
	combinedData.Status = false
	result, err := inst.componentsContainer.CombinedDataRepository().Create(userContext, combinedData)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.Create")
	}
	item, err := makeCombinedData(result)
	if err != nil {
		return nil, errutil.Wrap(err, "makeCombinedData")
	}
	return &pb.CombinedDataResponse{
		Data: item,
	}, nil
}

func (inst *RecommendService) CombinedDataApprove(ctx context.Context, req *pb.CombinedDataRequest) (*pb.CombinedDataResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "recommend-api", "update", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	combinedData, err := getCombinedData(req.Data)
	if err != nil {
		return nil, errutil.Wrap(err, "getCombinedData")
	}
	combinedDataItem, err := inst.componentsContainer.CombinedDataRepository().FindById(userContext, combinedData.CombinedDataId)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindById")
	}
	if combinedDataItem == nil {
		return nil, status.Errorf(codes.NotFound, "combinedDataItem not found")
	}
	if combinedDataItem.Status {
		return nil, status.Errorf(codes.FailedPrecondition, "These 2 tenants have been combined")
	}
	if !entityutil.ServiceOrAdminRole(userContext) {
		tenant, err := inst.componentsContainer.CustomerService().TenantFindById(userContext, combinedData.TenantId2)
		if err != nil {
			return nil, errutil.Wrap(err, "CustomerService.TenantFindById")
		}
		if tenant == nil {
			return nil, status.Errorf(codes.NotFound, "Tenant not found")
		}
		customerID, err := entityutil.GetUserId(userContext)
		if err != nil {
			return nil, errutil.Wrap(err, "entityutil.GetUserId")
		}
		if customerID != tenant.CustomerId {
			return nil, status.Errorf(codes.PermissionDenied, "PermissionDenied")
		}
	}
	combinedDataItem.Status = true
	countCategoryTenant1, err := inst.componentsContainer.SurveyService().GetCountCategory(userContext, combinedDataItem.TenantId1)
	if err != nil {
		return nil, errutil.Wrap(err, "SurveyService.GetCountCategory")
	}
	countCategoryTenant2, err := inst.componentsContainer.SurveyService().GetCountCategory(userContext, combinedDataItem.TenantId2)
	if err != nil {
		return nil, errutil.Wrap(err, "SurveyService.GetCountCategory")
	}
	combinedDataItem.NumberItem1 = countCategoryTenant1
	combinedDataItem.NumberItem2 = countCategoryTenant2
	combinedDataItem.SOnePart1 = countCategoryTenant1 * (countCategoryTenant1 + 5) / 2
	combinedDataItem.SOnePart2 = countCategoryTenant2 * (countCategoryTenant2 + 5) / 2
	combinedDataItem.NkOnePart1 = int32(math.Ceil(1.0/2 + math.Sqrt(float64(2*combinedDataItem.SOnePart1)+float64(1.0/4))))
	combinedDataItem.NkOnePart2 = int32(math.Ceil(1.0/2 + math.Sqrt(float64(2*combinedDataItem.SOnePart2)+float64(1.0/4))))
	combinedDataItem.STwoPart = countCategoryTenant1 * countCategoryTenant2
	combinedDataItem.SkTwoPart = int32(math.Ceil(1.0/2 + math.Sqrt(float64(2*combinedDataItem.STwoPart)+float64(1.0/4))))
	combinedDataItem.NkTwoPart = int32(math.Ceil(math.Sqrt(float64(combinedDataItem.SkTwoPart))))

	result, err := inst.componentsContainer.CombinedDataRepository().Update(userContext, combinedDataItem, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.Create")
	}
	item, err := makeCombinedData(result)
	if err != nil {
		return nil, errutil.Wrap(err, "makeCombinedData")
	}
	return &pb.CombinedDataResponse{
		Data: item,
	}, nil
}

func (inst *RecommendService) CombinedDataGetByTenantId(ctx context.Context, req *pb.StringRequest) (*pb.CombinedDataResponse, error) {
	userContext, err := inst.componentsContainer.AuthService().Authentication(ctx, req.Ctx.AccessToken, req.Ctx.DomainId, "@any", "@any", &grpcauth.AuthenOption{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, errutil.Message(err))
	}
	// if !entityutil.ServiceOrAdminRole(userContext) {
	// 	tenant, err := inst.componentsContainer.CustomerService().TenantFindById(userContext, req.Value)
	// 	if err != nil {
	// 		return nil, errutil.Wrap(err, "CustomerService.TenantFindById")
	// 	}
	// 	if tenant == nil {
	// 		return nil, status.Errorf(codes.NotFound, "Tenant not found")
	// 	}
	// 	customerID, err := entityutil.GetUserId(userContext)
	// 	if err != nil {
	// 		return nil, errutil.Wrap(err, "entityutil.GetUserId")
	// 	}
	// 	if customerID != tenant.CustomerId {
	// 		return nil, status.Errorf(codes.PermissionDenied, "PermissionDenied")
	// 	}
	// }
	combinedData, err := inst.componentsContainer.CombinedDataRepository().FindByTenantId(userContext, req.Value)
	if err != nil {
		return nil, errutil.Wrap(err, "CombinedDataRepository.FindByTenantId")
	}
	if combinedData == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Combined data not found")
	}
	item, err := makeCombinedData(combinedData)
	if err != nil {
		return nil, errutil.Wrap(err, "makeCombinedData")
	}
	return &pb.CombinedDataResponse{
		Data: item,
	}, nil
}
