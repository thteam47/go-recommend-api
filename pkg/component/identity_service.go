package component

import (
	"github.com/thteam47/common/entity"
)

type IdentityService interface {
	GetCountUsers(userContext entity.UserContext, tenantId string) (int32, error)
	GetUserById(userContext entity.UserContext, userId string) (*entity.User, error)
}
