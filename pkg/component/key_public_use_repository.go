package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type KeyPublicUseRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.KeyPublicUse, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.KeyPublicUse, error)
	FindById(userContext entity.UserContext, id string) (*models.KeyPublicUse, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.KeyPublicUse, error)
	Create(userContext entity.UserContext, user *models.KeyPublicUse) (*models.KeyPublicUse, error)
	Update(userContext entity.UserContext, data *models.KeyPublicUse, updateRequest *entity.UpdateRequest) (*models.KeyPublicUse, error)
	DeleteById(userContext entity.UserContext, id string) error
}
