package defaultcomponent

import (
	"context"
	"time"

	"github.com/thteam47/common-libs/confg"
	v1 "github.com/thteam47/common/api/customer-api"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
	"github.com/thteam47/go-recommend-api/util"
	"google.golang.org/grpc"
)

type CustomerService struct {
	config *CustomerServiceConfig
	client v1.CustomerServiceClient
}

type CustomerServiceConfig struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	AccessToken string        `mapstructure:"access_token"`
}

func NewCustomerServiceWithConfig(properties confg.Confg) (*CustomerService, error) {
	config := CustomerServiceConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}
	return NewCustomerService(&config)
}

func NewCustomerService(config *CustomerServiceConfig) (*CustomerService, error) {
	inst := &CustomerService{
		config: config,
	}
	conn, err := grpc.Dial(config.Address, grpc.WithInsecure())
	if err != nil {
		return nil, errutil.Wrapf(err, "grpc.Dial")
	}
	client := v1.NewCustomerServiceClient(conn)
	inst.client = client
	return inst, nil
}

func (inst *CustomerService) requestCtx(userContext entity.UserContext) *v1.Context {
	return &v1.Context{
		AccessToken: inst.config.AccessToken,
		DomainId:    userContext.DomainId(),
	}
}

func getTenant(item *v1.Tenant) (*models.Tenant, error) {
	if item == nil {
		return nil, nil
	}
	tenant := &models.Tenant{}
	err := util.FromMessage(item, tenant)
	if err != nil {
		return nil, errutil.Wrap(err, "FromMessage")
	}
	return tenant, nil
}

func (inst *CustomerService) TenantFindById(userContext entity.UserContext, tenantId string) (*models.Tenant, error) {
	result, err := inst.client.GetById(context.Background(), &v1.StringRequest{
		Ctx:   inst.requestCtx(userContext),
		Value: tenantId,
	})
	if err != nil {
		return nil, errutil.Wrapf(err, "client.GetById")
	}
	tenant, err := getTenant(result.Data)
	if err != nil {
		return nil, errutil.Wrapf(err, "getTenant")
	}
	return tenant, nil
}
