package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type Phase4TwoPartRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.Phase4TwoPart, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.Phase4TwoPart, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.Phase4TwoPart, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.Phase4TwoPart, error)
	Create(userContext entity.UserContext, user *models.Phase4TwoPart) (*models.Phase4TwoPart, error)
	Update(userContext entity.UserContext, data *models.Phase4TwoPart, updateRequest *entity.UpdateRequest) (*models.Phase4TwoPart, error)
	DeleteById(userContext entity.UserContext, id string) error
}
