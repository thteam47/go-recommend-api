package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type KeyPublicUserRepository struct {
	config         *KeyPublicUserRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type KeyPublicUserRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewKeyPublicUserRepositoryWithConfig(properties confg.Confg) (*KeyPublicUserRepository, error) {
	config := KeyPublicUserRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewKeyPublicUserRepository(&KeyPublicUserRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewKeyPublicUserRepository(config *KeyPublicUserRepositoryConfig) (*KeyPublicUserRepository, error) {
	inst := &KeyPublicUserRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.KeyPublicUser{},
		MongoIdField:       "Id",
		IdField:            "KeyPublicUserId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *KeyPublicUserRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.KeyPublicUser, error) {
	result := []models.KeyPublicUser{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *KeyPublicUserRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *KeyPublicUserRepository) FindById(userContext entity.UserContext, id string) (*models.KeyPublicUser, error) {
	result := &models.KeyPublicUser{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "KeyPublicUserId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *KeyPublicUserRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.KeyPublicUser, error) {
	result := &models.KeyPublicUser{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *KeyPublicUserRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.KeyPublicUser, error) {
	result := &models.KeyPublicUser{}
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
func (inst *KeyPublicUserRepository) Create(userContext entity.UserContext, data *models.KeyPublicUser) (*models.KeyPublicUser, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *KeyPublicUserRepository) Update(userContext entity.UserContext, data *models.KeyPublicUser, updateRequest *entity.UpdateRequest) (*models.KeyPublicUser, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "KeyPublicUserId", data.KeyPublicUserId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *KeyPublicUserRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "KeyPublicUserId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
