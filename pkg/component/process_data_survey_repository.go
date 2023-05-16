package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataSurveyRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataSurvey, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataSurvey, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.ProcessDataSurvey, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataSurvey, error)
	Create(userContext entity.UserContext, user *models.ProcessDataSurvey) (*models.ProcessDataSurvey, error)
	Update(userContext entity.UserContext, data *models.ProcessDataSurvey, updateRequest *entity.UpdateRequest) (*models.ProcessDataSurvey, error)
	DeleteById(userContext entity.UserContext, id string) error
}
