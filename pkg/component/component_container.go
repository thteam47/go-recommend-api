package component

import (
	grpcauth "github.com/thteam47/common/grpcutil"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/db"
)

type ComponentsContainer struct {
	combinedDataRepository       CombinedDataRepository
	authService                  *grpcauth.AuthInterceptor
	handler                      *db.Handler
	customerService              CustomerService
	surveyService                SurveyService
	identityService              IdentityService
	keyPublicUserRepository      KeyPublicUserRepository
	keyPublicUseRepository       KeyPublicUseRepository
	keyPublicItemRepository      KeyPublicItemRepository
	resultCardRepository         ResultCardRepository
	processDataSurveyRepository  ProcessDataSurveyRepository
	processDataSurvey2Repository ProcessDataSurvey2Repository
	phase3TwoPartRepository      Phase3TwoPartRepository
	phase4TwoPartRepository      Phase4TwoPartRepository
	processDataTotalRepository   ProcessDataTotalRepository
	processDataSumilorRepository ProcessDataSumilorRepository
	processDataRtbRepository     ProcessDataRtbRepository
}

func NewComponentsContainer(componentFactory ComponentFactory) (*ComponentsContainer, error) {
	inst := &ComponentsContainer{}

	var err error
	inst.authService = componentFactory.CreateAuthService()
	inst.combinedDataRepository, err = componentFactory.CreateCombinedDataRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateCombinedDataRepository")
	}
	inst.customerService, err = componentFactory.CreateCustomerService()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateCustomerService")
	}
	inst.surveyService, err = componentFactory.CreateSurveyService()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateSurveyService")
	}
	inst.identityService, err = componentFactory.CreateIdentityService()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateIdentityService")
	}
	inst.keyPublicUserRepository, err = componentFactory.CreateKeyPublicUserRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateKeyPublicUserRepository")
	}
	inst.keyPublicUseRepository, err = componentFactory.CreateKeyPublicUseRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateKeyPublicUseRepository")
	}
	inst.keyPublicItemRepository, err = componentFactory.CreateKeyPublicItemRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateKeyPublicItemRepository")
	}

	inst.resultCardRepository, err = componentFactory.CreateResultCardRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateResultCardRepository")
	}

	inst.processDataSurveyRepository, err = componentFactory.CreateProcessDataSurveyRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateProcessDataSurveyRepository")
	}
	inst.phase3TwoPartRepository, err = componentFactory.CreatePhase3TwoPartRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreatePhase3TwoPartRepository")
	}
	inst.phase4TwoPartRepository, err = componentFactory.CreatePhase4TwoPartRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreatePhase3TwoPartRepository")
	}

	inst.processDataSurvey2Repository, err = componentFactory.CreateProcessDataSurvey2Repository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateProcessDataSurveyRepository")
	}

	inst.processDataTotalRepository, err = componentFactory.CreateProcessDataTotalRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateProcessDataTotalRepository")
	}

	inst.processDataSumilorRepository, err = componentFactory.CreateProcessDataSumilorRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateProcessDataSumilorRepository")
	}

	inst.processDataRtbRepository, err = componentFactory.CreateProcessDataRtbRepository()
	if err != nil {
		return nil, errutil.Wrap(err, "CreateProcessDataRtbRepository")
	}

	return inst, nil
}

func (inst *ComponentsContainer) AuthService() *grpcauth.AuthInterceptor {
	return inst.authService
}

func (inst *ComponentsContainer) CombinedDataRepository() CombinedDataRepository {
	return inst.combinedDataRepository
}

func (inst *ComponentsContainer) CustomerService() CustomerService {
	return inst.customerService
}

func (inst *ComponentsContainer) SurveyService() SurveyService {
	return inst.surveyService
}

func (inst *ComponentsContainer) IdentityService() IdentityService {
	return inst.identityService
}

func (inst *ComponentsContainer) KeyPublicItemRepository() KeyPublicItemRepository {
	return inst.keyPublicItemRepository
}

func (inst *ComponentsContainer) KeyPublicUserRepository() KeyPublicUserRepository {
	return inst.keyPublicUserRepository
}

func (inst *ComponentsContainer) KeyPublicUseRepository() KeyPublicUseRepository {
	return inst.keyPublicUseRepository
}

func (inst *ComponentsContainer) ResultCardRepository() ResultCardRepository {
	return inst.resultCardRepository
}

func (inst *ComponentsContainer) ProcessDataSurveyRepository() ProcessDataSurveyRepository {
	return inst.processDataSurveyRepository
}

func (inst *ComponentsContainer) ProcessDataSurvey2Repository() ProcessDataSurvey2Repository {
	return inst.processDataSurvey2Repository
}

func (inst *ComponentsContainer) Phase3TwoPartRepository() Phase3TwoPartRepository {
	return inst.phase3TwoPartRepository
}

func (inst *ComponentsContainer) Phase4TwoPartRepository() Phase4TwoPartRepository {
	return inst.phase4TwoPartRepository
}

func (inst *ComponentsContainer) ProcessDataTotalRepository() ProcessDataTotalRepository {
	return inst.processDataTotalRepository
}

func (inst *ComponentsContainer) ProcessDataSumilorRepository() ProcessDataSumilorRepository {
	return inst.processDataSumilorRepository
}

func (inst *ComponentsContainer) ProcessDataRtbRepository() ProcessDataRtbRepository {
	return inst.processDataRtbRepository
}

var errorCodeBadRequest = 400
