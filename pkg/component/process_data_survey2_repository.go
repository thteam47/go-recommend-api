package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataSurvey2Repository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataSurvey2, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataSurvey2, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.ProcessDataSurvey2, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataSurvey2, error)
	Create(userContext entity.UserContext, user *models.ProcessDataSurvey2) (*models.ProcessDataSurvey2, error)
	Update(userContext entity.UserContext, data *models.ProcessDataSurvey2, updateRequest *entity.UpdateRequest) (*models.ProcessDataSurvey2, error)
	DeleteById(userContext entity.UserContext, id string) error
}
