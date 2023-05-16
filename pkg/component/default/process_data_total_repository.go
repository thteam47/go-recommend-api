package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataTotalRepository struct {
	config         *ProcessDataTotalRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type ProcessDataTotalRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewProcessDataTotalRepositoryWithConfig(properties confg.Confg) (*ProcessDataTotalRepository, error) {
	config := ProcessDataTotalRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewProcessDataTotalRepository(&ProcessDataTotalRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewProcessDataTotalRepository(config *ProcessDataTotalRepositoryConfig) (*ProcessDataTotalRepository, error) {
	inst := &ProcessDataTotalRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.ProcessDataTotal{},
		MongoIdField:       "Id",
		IdField:            "ProcessDataTotalId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *ProcessDataTotalRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataTotal, error) {
	result := []models.ProcessDataTotal{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *ProcessDataTotalRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *ProcessDataTotalRepository) FindById(userContext entity.UserContext, id string) (*models.ProcessDataTotal, error) {
	result := &models.ProcessDataTotal{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "ProcessDataTotalId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataTotalRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataTotal, error) {
	result := &models.ProcessDataTotal{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataTotalRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataTotal, error) {
	result := &models.ProcessDataTotal{}
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
func (inst *ProcessDataTotalRepository) Create(userContext entity.UserContext, data *models.ProcessDataTotal) (*models.ProcessDataTotal, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *ProcessDataTotalRepository) CreateAndUpdate(userContext entity.UserContext, data *models.ProcessDataTotal) error {
	result := &models.ProcessDataTotal{}
	err := inst.baseRepository.FindOneByAttributes(userContext, map[string]interface{}{
		"DomainId":     data.DomainId,
		"PositionItem": data.PositionItem,
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
		err = inst.baseRepository.UpdateOneByAttribute(userContext, "ProcessDataTotalId", result.ProcessDataTotalId, data, nil, &mongorepository.UpdateOptions{
			ExcludedProperties: excludedProperties,
			InsertIfNotExisted: true,
		})
		if err != nil {
			return errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
		}
	}

	return nil
}

func (inst *ProcessDataTotalRepository) Update(userContext entity.UserContext, data *models.ProcessDataTotal, updateRequest *entity.UpdateRequest) (*models.ProcessDataTotal, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "ProcessDataTotalId", data.ProcessDataTotalId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *ProcessDataTotalRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "ProcessDataTotalId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
