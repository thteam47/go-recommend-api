package component

import (
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type CustomerService interface {
	TenantFindById(userContext entity.UserContext, tenantId string) (*models.Tenant, error)
}
