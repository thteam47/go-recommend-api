package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type Phase3TwoPartRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.Phase3TwoPart, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.Phase3TwoPart, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.Phase3TwoPart, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.Phase3TwoPart, error)
	Create(userContext entity.UserContext, user *models.Phase3TwoPart) (*models.Phase3TwoPart, error)
	Update(userContext entity.UserContext, data *models.Phase3TwoPart, updateRequest *entity.UpdateRequest) (*models.Phase3TwoPart, error)
	DeleteById(userContext entity.UserContext, id string) error
}
