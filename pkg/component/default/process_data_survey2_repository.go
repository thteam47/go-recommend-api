package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataSurvey2Repository struct {
	config         *ProcessDataSurvey2RepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type ProcessDataSurvey2RepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewProcessDataSurvey2RepositoryWithConfig(properties confg.Confg) (*ProcessDataSurvey2Repository, error) {
	config := ProcessDataSurvey2RepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewProcessDataSurvey2Repository(&ProcessDataSurvey2RepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewProcessDataSurvey2Repository(config *ProcessDataSurvey2RepositoryConfig) (*ProcessDataSurvey2Repository, error) {
	inst := &ProcessDataSurvey2Repository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.ProcessDataSurvey2{},
		MongoIdField:       "Id",
		IdField:            "ProcessDataSurveyId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *ProcessDataSurvey2Repository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataSurvey2, error) {
	result := []models.ProcessDataSurvey2{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *ProcessDataSurvey2Repository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *ProcessDataSurvey2Repository) FindById(userContext entity.UserContext, id string) (*models.ProcessDataSurvey2, error) {
	result := &models.ProcessDataSurvey2{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "ProcessDataSurvey2Id", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataSurvey2Repository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataSurvey2, error) {
	result := &models.ProcessDataSurvey2{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataSurvey2Repository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataSurvey2, error) {
	result := &models.ProcessDataSurvey2{}
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
func (inst *ProcessDataSurvey2Repository) Create(userContext entity.UserContext, data *models.ProcessDataSurvey2) (*models.ProcessDataSurvey2, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *ProcessDataSurvey2Repository) Update(userContext entity.UserContext, data *models.ProcessDataSurvey2, updateRequest *entity.UpdateRequest) (*models.ProcessDataSurvey2, error) {
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
func (inst *ProcessDataSurvey2Repository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "ProcessDataSurvey2Id", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
