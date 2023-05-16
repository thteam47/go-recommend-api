package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type KeyPublicUserRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.KeyPublicUser, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.KeyPublicUser, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.KeyPublicUser, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.KeyPublicUser, error)
	Create(userContext entity.UserContext, user *models.KeyPublicUser) (*models.KeyPublicUser, error)
	Update(userContext entity.UserContext, data *models.KeyPublicUser, updateRequest *entity.UpdateRequest) (*models.KeyPublicUser, error)
	DeleteById(userContext entity.UserContext, id string) error
}
