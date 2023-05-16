package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/common/handler"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/component"
)

type ComponentFactory struct {
	properties confg.Confg
	handle     *handler.Handler
}

func NewComponentFactory(properties confg.Confg, handle *handler.Handler) (*ComponentFactory, error) {
	inst := &ComponentFactory{
		properties: properties,
		handle:     handle,
	}

	return inst, nil
}

func (inst *ComponentFactory) CreateAuthService() *grpcauth.AuthInterceptor {
	authService := grpcauth.NewAuthInterceptor(inst.handle)
	return authService
}

func (inst *ComponentFactory) CreateCombinedDataRepository() (component.CombinedDataRepository, error) {
	combinedDataRepository, err := NewCombinedDataRepositoryWithConfig(inst.properties.Sub("combined-data-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewCombinedDataRepositoryWithConfig")
	}
	return combinedDataRepository, nil
}

func (inst *ComponentFactory) CreateCustomerService() (component.CustomerService, error) {
	customerService, err := NewCustomerServiceWithConfig(inst.properties.Sub("customer-service"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewCustomerServiceWithConfig")
	}
	return customerService, nil
}

func (inst *ComponentFactory) CreateSurveyService() (component.SurveyService, error) {
	surveyService, err := NewSurveyServiceWithConfig(inst.properties.Sub("survey-service"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewCustomerServiceWithConfig")
	}
	return surveyService, nil
}

func (inst *ComponentFactory) CreateIdentityService() (component.IdentityService, error) {
	identityService, err := NewIdentityServiceWithConfig(inst.properties.Sub("identity-service"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewCustomerServiceWithConfig")
	}
	return identityService, nil
}

func (inst *ComponentFactory) CreateKeyPublicItemRepository() (component.KeyPublicItemRepository, error) {
	keyPublicItemRepository, err := NewKeyPublicItemRepositoryWithConfig(inst.properties.Sub("key-public-item-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewKeyPublicItemRepositoryWithConfig")
	}
	return keyPublicItemRepository, nil
}

func (inst *ComponentFactory) CreateKeyPublicUserRepository() (component.KeyPublicUserRepository, error) {
	keyPublicUserRepository, err := NewKeyPublicUserRepositoryWithConfig(inst.properties.Sub("key-public-user-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewKeyPublicUserRepositoryWithConfig")
	}
	return keyPublicUserRepository, nil
}

func (inst *ComponentFactory) CreateKeyPublicUseRepository() (component.KeyPublicUseRepository, error) {
	keyPublicUseRepository, err := NewKeyPublicUseRepositoryWithConfig(inst.properties.Sub("key-public-use-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewKeyPublicUseRepositoryWithConfig")
	}
	return keyPublicUseRepository, nil
}

func (inst *ComponentFactory) CreateResultCardRepository() (component.ResultCardRepository, error) {
	resultCardRepository, err := NewResultCardRepositoryWithConfig(inst.properties.Sub("result-card-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewResultCardRepositoryWithConfig")
	}
	return resultCardRepository, nil
}

func (inst *ComponentFactory) CreateProcessDataSurveyRepository() (component.ProcessDataSurveyRepository, error) {
	processDataSurveyRepository, err := NewProcessDataSurveyRepositoryWithConfig(inst.properties.Sub("process-data-survey-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewProcessDataSurveyRepositoryWithConfig")
	}
	return processDataSurveyRepository, nil
}

func (inst *ComponentFactory) CreateProcessDataTotalRepository() (component.ProcessDataTotalRepository, error) {
	processDataTotalRepository, err := NewProcessDataTotalRepositoryWithConfig(inst.properties.Sub("process-data-total-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewProcessDataTotalRepositoryWithConfig")
	}
	return processDataTotalRepository, nil
}

func (inst *ComponentFactory) CreateProcessDataSumilorRepository() (component.ProcessDataSumilorRepository, error) {
	processDataSumilorRepository, err := NewProcessDataSumilorRepositoryWithConfig(inst.properties.Sub("process-data-sumilor-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewProcessDataSumilorRepositoryWithConfig")
	}
	return processDataSumilorRepository, nil
}

func (inst *ComponentFactory) CreateProcessDataRtbRepository() (component.ProcessDataRtbRepository, error) {
	processDataRtbRepository, err := NewProcessDataRtbRepositoryWithConfig(inst.properties.Sub("process-data-rtb-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewProcessDataRtbRepositoryWithConfig")
	}
	return processDataRtbRepository, nil
}

func (inst *ComponentFactory) CreateProcessDataSurvey2Repository() (component.ProcessDataSurvey2Repository, error) {
	processDataSurvey2Repository, err := NewProcessDataSurvey2RepositoryWithConfig(inst.properties.Sub("process-data-survey2-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewProcessDataSurveyRepositoryWithConfig")
	}
	return processDataSurvey2Repository, nil
}

func (inst *ComponentFactory) CreatePhase3TwoPartRepository() (component.Phase3TwoPartRepository, error) {
	phase3TwoPartRepository, err := NewPhase3TwoPartRepositoryWithConfig(inst.properties.Sub("phase3-two-part-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewPhase3TwoPartRepositoryWithConfig")
	}
	return phase3TwoPartRepository, nil
}

func (inst *ComponentFactory) CreatePhase4TwoPartRepository() (component.Phase4TwoPartRepository, error) {
	phase4TwoPartRepository, err := NewPhase4TwoPartRepositoryWithConfig(inst.properties.Sub("phase4-two-part-repository"))
	if err != nil {
		return nil, errutil.Wrapf(err, "NewPhase4TwoPartRepositoryWithConfig")
	}
	return phase4TwoPartRepository, nil
}
