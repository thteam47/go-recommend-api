package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataSumilorRepository struct {
	config         *ProcessDataSumilorRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type ProcessDataSumilorRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewProcessDataSumilorRepositoryWithConfig(properties confg.Confg) (*ProcessDataSumilorRepository, error) {
	config := ProcessDataSumilorRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewProcessDataSumilorRepository(&ProcessDataSumilorRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewProcessDataSumilorRepository(config *ProcessDataSumilorRepositoryConfig) (*ProcessDataSumilorRepository, error) {
	inst := &ProcessDataSumilorRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.ProcessDataSumilor{},
		MongoIdField:       "Id",
		IdField:            "ProcessDataSumilorId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *ProcessDataSumilorRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataSumilor, error) {
	result := []models.ProcessDataSumilor{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *ProcessDataSumilorRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *ProcessDataSumilorRepository) FindById(userContext entity.UserContext, id string) (*models.ProcessDataSumilor, error) {
	result := &models.ProcessDataSumilor{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "ProcessDataSumilorId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataSumilorRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataSumilor, error) {
	result := &models.ProcessDataSumilor{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataSumilorRepository) FindByOneByFindRequest(userContext entity.UserContext, findRequest *entity.FindRequest) (*models.ProcessDataSumilor, error) {
	result := &models.ProcessDataSumilor{}
	err := inst.baseRepository.FindOneByFindRequest(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByFindRequest")
	}
	return result, nil
}

func (inst *ProcessDataSumilorRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataSumilor, error) {
	result := &models.ProcessDataSumilor{}
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
func (inst *ProcessDataSumilorRepository) Create(userContext entity.UserContext, data *models.ProcessDataSumilor) (*models.ProcessDataSumilor, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *ProcessDataSumilorRepository) CreateAndUpdate(userContext entity.UserContext, data *models.ProcessDataSumilor) error {
	result := &models.ProcessDataSumilor{}
	err := inst.baseRepository.FindOneByAttributes(userContext, map[string]interface{}{
		"DomainId":     data.DomainId,
		"PositionItemOriginal1": data.PositionItemOriginal1,
		"PositionItemOriginal2": data.PositionItemOriginal2,
	}, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	if result == nil {
		err := inst.baseRepository.Create(userContext, data, nil)
		if err != nil {
			return errutil.Wrap(err, "baseRepository.Create")
		}
	} else {
		excludedProperties := []string{
			"CreatedTime",
		}
		err = inst.baseRepository.UpdateOneByAttribute(userContext, "ProcessDataSumilorId", result.ProcessDataSumilorId, data, nil, &mongorepository.UpdateOptions{
			ExcludedProperties: excludedProperties,
			InsertIfNotExisted: true,
		})
		if err != nil {
			return errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
		}
	}

	return nil
}

func (inst *ProcessDataSumilorRepository) Update(userContext entity.UserContext, data *models.ProcessDataSumilor, updateRequest *entity.UpdateRequest) (*models.ProcessDataSumilor, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "ProcessDataSumilorId", data.ProcessDataSumilorId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *ProcessDataSumilorRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "ProcessDataSumilorId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
