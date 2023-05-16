package defaultcomponent

import (
	"context"
	"time"

	"github.com/thteam47/common-libs/confg"
	v1 "github.com/thteam47/common/api/survey-api"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
	"github.com/thteam47/go-recommend-api/util"
	"google.golang.org/grpc"
)

type SurveyService struct {
	config *SurveyServiceConfig
	client v1.SurveyServiceClient
}

type SurveyServiceConfig struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	AccessToken string        `mapstructure:"access_token"`
}

func NewSurveyServiceWithConfig(properties confg.Confg) (*SurveyService, error) {
	config := SurveyServiceConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}
	return NewSurveyService(&config)
}

func NewSurveyService(config *SurveyServiceConfig) (*SurveyService, error) {
	inst := &SurveyService{
		config: config,
	}
	conn, err := grpc.Dial(config.Address, grpc.WithInsecure())
	if err != nil {
		return nil, errutil.Wrapf(err, "grpc.Dial")
	}
	client := v1.NewSurveyServiceClient(conn)
	inst.client = client
	return inst, nil
}

func (inst *SurveyService) requestCtx(userContext entity.UserContext) *v1.Context {
	return &v1.Context{
		AccessToken: inst.config.AccessToken,
		DomainId:    userContext.DomainId(),
	}
}

func getCategory(item *v1.Category) (*models.Category, error) {
	if item == nil {
		return nil, nil
	}
	category := &models.Category{}
	err := util.FromMessage(item, category)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return category, nil
}

func (inst *SurveyService) GetCountCategory(userContext entity.UserContext, tenantId string) (int32, error) {
	result, err := inst.client.GetAllCategory(context.Background(), &v1.ListRequest{
		Ctx:   inst.requestCtx(userContext),
		Limit: -10,
		Filters: []*v1.ListRequest_Filter{
			&v1.ListRequest_Filter{
				Key: "DomainId",
				Value: tenantId,
				Operator: entity.FindRequestFilterOperatorEqualTo,
			},
		},
	})
	if err != nil {
		return 0, errutil.Wrapf(err, "client.GetById")
	}
	return result.Total, nil
}
