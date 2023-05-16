package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataSumilorRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataSumilor, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataSumilor, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.ProcessDataSumilor, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataSumilor, error)
	Create(userContext entity.UserContext, user *models.ProcessDataSumilor) (*models.ProcessDataSumilor, error)
	CreateAndUpdate(userContext entity.UserContext, data *models.ProcessDataSumilor) error
	Update(userContext entity.UserContext, data *models.ProcessDataSumilor, updateRequest *entity.UpdateRequest) (*models.ProcessDataSumilor, error)
	DeleteById(userContext entity.UserContext, id string) error
	FindByOneByFindRequest(userContext entity.UserContext, findRequest *entity.FindRequest) (*models.ProcessDataSumilor, error)
}
