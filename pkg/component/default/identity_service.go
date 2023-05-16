package defaultcomponent

import (
	"context"
	"time"

	"github.com/thteam47/common-libs/confg"
	v1 "github.com/thteam47/common/api/identity-api"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/util"
	"google.golang.org/grpc"
)

type IdentityService struct {
	config *IdentityServiceConfig
	client v1.IdentityServiceClient
}

type IdentityServiceConfig struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	AccessToken string        `mapstructure:"access_token"`
}

func NewIdentityServiceWithConfig(properties confg.Confg) (*IdentityService, error) {
	config := IdentityServiceConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}
	return NewIdentityService(&config)
}

func NewIdentityService(config *IdentityServiceConfig) (*IdentityService, error) {
	inst := &IdentityService{
		config: config,
	}
	conn, err := grpc.Dial(config.Address, grpc.WithInsecure())
	if err != nil {
		return nil, errutil.Wrapf(err, "grpc.Dial")
	}
	client := v1.NewIdentityServiceClient(conn)
	inst.client = client
	return inst, nil
}

func (inst *IdentityService) requestCtx(userContext entity.UserContext) *v1.Context {
	return &v1.Context{
		AccessToken: inst.config.AccessToken,
		DomainId:    userContext.DomainId(),
	}
}

func getUser(item *v1.User) (*entity.User, error) {
	if item == nil {
		return nil, nil
	}
	user := &entity.User{}
	err := util.FromMessage(item, user)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return user, nil
}

func (inst *IdentityService) GetCountUsers(userContext entity.UserContext, tenantId string) (int32, error) {
	result, err := inst.client.GetAll(context.Background(), &v1.ListRequest{
		Ctx:   inst.requestCtx(userContext),
		Limit: -10,
		Filters: []*v1.ListRequest_Filter{
			&v1.ListRequest_Filter{
				Key:      "DomainId",
				Value:    tenantId,
				Operator: entity.FindRequestFilterOperatorEqualTo,
			},
		},
	})
	if err != nil {
		return 0, errutil.Wrapf(err, "client.GetById")
	}
	return result.Total, nil
}

func (inst *IdentityService) GetUserById(userContext entity.UserContext, userId string) (*entity.User, error) {
	result, err := inst.client.GetById(context.Background(), &v1.StringRequest{
		Ctx:   inst.requestCtx(userContext),
		Value: userId,
	})
	if err != nil {
		return nil, errutil.Wrapf(err, "client.GetById")
	}
	user, err := getUser(result.Data)
	if err != nil {
		return nil, errutil.Wrapf(err, "getUser")
	}
	return user, nil
}