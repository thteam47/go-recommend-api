package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataTotalRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataTotal, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataTotal, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.ProcessDataTotal, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataTotal, error)
	Create(userContext entity.UserContext, user *models.ProcessDataTotal) (*models.ProcessDataTotal, error)
	CreateAndUpdate(userContext entity.UserContext, data *models.ProcessDataTotal) error
	Update(userContext entity.UserContext, data *models.ProcessDataTotal, updateRequest *entity.UpdateRequest) (*models.ProcessDataTotal, error)
	DeleteById(userContext entity.UserContext, id string) error
}
