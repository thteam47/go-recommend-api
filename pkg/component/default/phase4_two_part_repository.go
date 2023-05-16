package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type Phase4TwoPartRepository struct {
	config         *Phase4TwoPartRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type Phase4TwoPartRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewPhase4TwoPartRepositoryWithConfig(properties confg.Confg) (*Phase4TwoPartRepository, error) {
	config := Phase4TwoPartRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewPhase4TwoPartRepository(&Phase4TwoPartRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewPhase4TwoPartRepository(config *Phase4TwoPartRepositoryConfig) (*Phase4TwoPartRepository, error) {
	inst := &Phase4TwoPartRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.Phase4TwoPart{},
		MongoIdField:       "Id",
		IdField:            "Phase4TwoPartId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *Phase4TwoPartRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.Phase4TwoPart, error) {
	result := []models.Phase4TwoPart{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *Phase4TwoPartRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *Phase4TwoPartRepository) FindById(userContext entity.UserContext, id string) (*models.Phase4TwoPart, error) {
	result := &models.Phase4TwoPart{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "Phase4TwoPartId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *Phase4TwoPartRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.Phase4TwoPart, error) {
	result := &models.Phase4TwoPart{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *Phase4TwoPartRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.Phase4TwoPart, error) {
	result := &models.Phase4TwoPart{}
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
func (inst *Phase4TwoPartRepository) Create(userContext entity.UserContext, data *models.Phase4TwoPart) (*models.Phase4TwoPart, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *Phase4TwoPartRepository) Update(userContext entity.UserContext, data *models.Phase4TwoPart, updateRequest *entity.UpdateRequest) (*models.Phase4TwoPart, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "Phase4TwoPartId", data.Phase4TwoPartId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *Phase4TwoPartRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "Phase4TwoPartId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
