package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataSurveyRepository struct {
	config         *ProcessDataSurveyRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type ProcessDataSurveyRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewProcessDataSurveyRepositoryWithConfig(properties confg.Confg) (*ProcessDataSurveyRepository, error) {
	config := ProcessDataSurveyRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewProcessDataSurveyRepository(&ProcessDataSurveyRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewProcessDataSurveyRepository(config *ProcessDataSurveyRepositoryConfig) (*ProcessDataSurveyRepository, error) {
	inst := &ProcessDataSurveyRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.ProcessDataSurvey{},
		MongoIdField:       "Id",
		IdField:            "ProcessDataSurveyId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *ProcessDataSurveyRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataSurvey, error) {
	result := []models.ProcessDataSurvey{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *ProcessDataSurveyRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *ProcessDataSurveyRepository) FindById(userContext entity.UserContext, id string) (*models.ProcessDataSurvey, error) {
	result := &models.ProcessDataSurvey{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "ProcessDataSurveyId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataSurveyRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataSurvey, error) {
	result := &models.ProcessDataSurvey{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataSurveyRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataSurvey, error) {
	result := &models.ProcessDataSurvey{}
	err := inst.baseRepository.FindOneByFindRequest(userContext, &entity.FindRequest{
		Filters: []entity.FindRequestFilter{
			entity.FindRequestFilter{
				Operator: entity.FindRequestFilterOperatorOr,
				SubFilters: []entity.FindRequestFilter{
					entity.FindRequestFilter{
						Key:      "TenantId1",
						Operator: entity.FindRequestFilterOperatorEqualTo,
						Value:    tenantId,
					},
					entity.FindRequestFilter{
						Key:      "TenantId2",
						Operator: entity.FindRequestFilterOperatorEqualTo,
						Value:    tenantId,
					},
				},
			},
		},
	}, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByFindRequest")
	}
	return result, nil
}
func (inst *ProcessDataSurveyRepository) Create(userContext entity.UserContext, data *models.ProcessDataSurvey) (*models.ProcessDataSurvey, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *ProcessDataSurveyRepository) Update(userContext entity.UserContext, data *models.ProcessDataSurvey, updateRequest *entity.UpdateRequest) (*models.ProcessDataSurvey, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "ProcessDataSurveyId", data.ProcessDataSurveyId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *ProcessDataSurveyRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "ProcessDataSurveyId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
