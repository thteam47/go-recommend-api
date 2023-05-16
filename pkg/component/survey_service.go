package component

import (
	"github.com/thteam47/common/entity"
)

type SurveyService interface {
	GetCountCategory(userContext entity.UserContext, tenantId string) (int32, error)
}
