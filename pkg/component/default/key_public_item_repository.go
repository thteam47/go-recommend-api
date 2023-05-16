package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type KeyPublicItemRepository struct {
	config         *KeyPublicItemRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type KeyPublicItemRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewKeyPublicItemRepositoryWithConfig(properties confg.Confg) (*KeyPublicItemRepository, error) {
	config := KeyPublicItemRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewKeyPublicItemRepository(&KeyPublicItemRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewKeyPublicItemRepository(config *KeyPublicItemRepositoryConfig) (*KeyPublicItemRepository, error) {
	inst := &KeyPublicItemRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.KeyPublicItem{},
		MongoIdField:       "Id",
		IdField:            "KeyPublicItemId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *KeyPublicItemRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.KeyPublicItem, error) {
	result := []models.KeyPublicItem{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *KeyPublicItemRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *KeyPublicItemRepository) FindById(userContext entity.UserContext, id string) (*models.KeyPublicItem, error) {
	result := &models.KeyPublicItem{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "KeyPublicItemId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *KeyPublicItemRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.KeyPublicItem, error) {
	result := &models.KeyPublicItem{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *KeyPublicItemRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.KeyPublicItem, error) {
	result := &models.KeyPublicItem{}
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
func (inst *KeyPublicItemRepository) Create(userContext entity.UserContext, data *models.KeyPublicItem) (*models.KeyPublicItem, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *KeyPublicItemRepository) Update(userContext entity.UserContext, data *models.KeyPublicItem, updateRequest *entity.UpdateRequest) (*models.KeyPublicItem, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "KeyPublicItemId", data.KeyPublicItemId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *KeyPublicItemRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "KeyPublicItemId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
