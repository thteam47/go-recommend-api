package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataRtbRepository interface {
	FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataRtb, error)
	FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataRtb, error)
	Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error)
	FindById(userContext entity.UserContext, id string) (*models.ProcessDataRtb, error)
	FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataRtb, error)
	Create(userContext entity.UserContext, user *models.ProcessDataRtb) (*models.ProcessDataRtb, error)
	CreateAndUpdate(userContext entity.UserContext, data *models.ProcessDataRtb) error
	Update(userContext entity.UserContext, data *models.ProcessDataRtb, updateRequest *entity.UpdateRequest) (*models.ProcessDataRtb, error)
	DeleteById(userContext entity.UserContext, id string) error
	FindByOneByFindRequest(userContext entity.UserContext, findRequest *entity.FindRequest) (*models.ProcessDataRtb, error)
}
