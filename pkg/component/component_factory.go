package component

import grpcauth "github.com/thteam47/common/grpcutil"

type ComponentFactory interface {
	CreateAuthService() *grpcauth.AuthInterceptor
	CreateCombinedDataRepository() (CombinedDataRepository, error)
	CreateCustomerService() (CustomerService, error)
	CreateSurveyService() (SurveyService, error)
	CreateIdentityService() (IdentityService, error)
	CreateKeyPublicUseRepository() (KeyPublicUseRepository, error)
	CreateKeyPublicUserRepository() (KeyPublicUserRepository, error)
	CreateKeyPublicItemRepository() (KeyPublicItemRepository, error)
	CreateResultCardRepository() (ResultCardRepository, error)
	CreateProcessDataSurveyRepository() (ProcessDataSurveyRepository, error)
	CreateProcessDataTotalRepository() (ProcessDataTotalRepository, error)
	CreatePhase3TwoPartRepository() (Phase3TwoPartRepository, error)
	CreatePhase4TwoPartRepository() (Phase4TwoPartRepository, error)
	CreateProcessDataSurvey2Repository() (ProcessDataSurvey2Repository, error)
	CreateProcessDataSumilorRepository() (ProcessDataSumilorRepository, error)
	CreateProcessDataRtbRepository() (ProcessDataRtbRepository, error)
}
