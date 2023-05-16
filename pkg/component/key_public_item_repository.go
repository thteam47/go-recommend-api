package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type KeyPublicItemRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.KeyPublicItem, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.KeyPublicItem, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.KeyPublicItem, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.KeyPublicItem, error)
	Create(userContext entity.UserContext, user *models.KeyPublicItem) (*models.KeyPublicItem, error)
	Update(userContext entity.UserContext, data *models.KeyPublicItem, updateRequest *entity.UpdateRequest) (*models.KeyPublicItem, error)
	DeleteById(userContext entity.UserContext, id string) error
}
