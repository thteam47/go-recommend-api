package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type CombinedDataRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.CombinedData, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.CombinedData, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.CombinedData, error)
	Create(userContext entity.UserContext, user *models.CombinedData) (*models.CombinedData, error)
	Update(userContext entity.UserContext, data *models.CombinedData, updateRequest *entity.UpdateRequest) (*models.CombinedData, error)
	DeleteById(userContext entity.UserContext, id string) error
}
