package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type KeyPublicUseRepository struct {
	config         *KeyPublicUseRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type KeyPublicUseRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewKeyPublicUseRepositoryWithConfig(properties confg.Confg) (*KeyPublicUseRepository, error) {
	config := KeyPublicUseRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewKeyPublicUseRepository(&KeyPublicUseRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewKeyPublicUseRepository(config *KeyPublicUseRepositoryConfig) (*KeyPublicUseRepository, error) {
	inst := &KeyPublicUseRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.KeyPublicUse{},
		MongoIdField:       "Id",
		IdField:            "KeyPublicUseId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *KeyPublicUseRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.KeyPublicUse, error) {
	result := []models.KeyPublicUse{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *KeyPublicUseRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *KeyPublicUseRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.KeyPublicUse, error) {
	result := &models.KeyPublicUse{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *KeyPublicUseRepository) FindById(userContext entity.UserContext, id string) (*models.KeyPublicUse, error) {
	result := &models.KeyPublicUse{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "KeyPublicUseId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *KeyPublicUseRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.KeyPublicUse, error) {
	result := &models.KeyPublicUse{}
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
func (inst *KeyPublicUseRepository) Create(userContext entity.UserContext, data *models.KeyPublicUse) (*models.KeyPublicUse, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *KeyPublicUseRepository) Update(userContext entity.UserContext, data *models.KeyPublicUse, updateRequest *entity.UpdateRequest) (*models.KeyPublicUse, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "KeyPublicUseId", data.KeyPublicUseId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *KeyPublicUseRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "KeyPublicUseId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
