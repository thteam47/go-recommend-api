package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type CombinedDataRepository struct {
	config         *CombinedDataRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type CombinedDataRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewCombinedDataRepositoryWithConfig(properties confg.Confg) (*CombinedDataRepository, error) {
	config := CombinedDataRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewCombinedDataRepository(&CombinedDataRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewCombinedDataRepository(config *CombinedDataRepositoryConfig) (*CombinedDataRepository, error) {
	inst := &CombinedDataRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.CombinedData{},
		MongoIdField:       "Id",
		IdField:            "CombinedDataId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *CombinedDataRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.CombinedData, error) {
	result := []models.CombinedData{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *CombinedDataRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *CombinedDataRepository) FindById(userContext entity.UserContext, id string) (*models.CombinedData, error) {
	result := &models.CombinedData{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "CombinedDataId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *CombinedDataRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.CombinedData, error) {
	result := &models.CombinedData{}
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
func (inst *CombinedDataRepository) Create(userContext entity.UserContext, data *models.CombinedData) (*models.CombinedData, error) {
	if data.Meta == nil {
		data.Meta = map[string]string{}
	}
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *CombinedDataRepository) Update(userContext entity.UserContext, data *models.CombinedData, updateRequest *entity.UpdateRequest) (*models.CombinedData, error) {
	if data.Meta == nil {
		data.Meta = map[string]string{}
	}
	excludedProperties := []string{
		"CreatedTime", "TenantId1", "TenantId2",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "CombinedDataId", data.CombinedDataId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *CombinedDataRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "CombinedDataId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
