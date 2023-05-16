package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ResultCardRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ResultCard, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ResultCard, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.ResultCard, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ResultCard, error)
	Create(userContext entity.UserContext, user *models.ResultCard) (*models.ResultCard, error)
	Update(userContext entity.UserContext, data *models.ResultCard, updateRequest *entity.UpdateRequest) (*models.ResultCard, error)
	DeleteById(userContext entity.UserContext, id string) error
}
