package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type Phase3TwoPartRepository struct {
	config         *Phase3TwoPartRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type Phase3TwoPartRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewPhase3TwoPartRepositoryWithConfig(properties confg.Confg) (*Phase3TwoPartRepository, error) {
	config := Phase3TwoPartRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewPhase3TwoPartRepository(&Phase3TwoPartRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewPhase3TwoPartRepository(config *Phase3TwoPartRepositoryConfig) (*Phase3TwoPartRepository, error) {
	inst := &Phase3TwoPartRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.Phase3TwoPart{},
		MongoIdField:       "Id",
		IdField:            "Phase3TwoPartId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *Phase3TwoPartRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.Phase3TwoPart, error) {
	result := []models.Phase3TwoPart{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *Phase3TwoPartRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *Phase3TwoPartRepository) FindById(userContext entity.UserContext, id string) (*models.Phase3TwoPart, error) {
	result := &models.Phase3TwoPart{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "Phase3TwoPartId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *Phase3TwoPartRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.Phase3TwoPart, error) {
	result := &models.Phase3TwoPart{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *Phase3TwoPartRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.Phase3TwoPart, error) {
	result := &models.Phase3TwoPart{}
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
func (inst *Phase3TwoPartRepository) Create(userContext entity.UserContext, data *models.Phase3TwoPart) (*models.Phase3TwoPart, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *Phase3TwoPartRepository) Update(userContext entity.UserContext, data *models.Phase3TwoPart, updateRequest *entity.UpdateRequest) (*models.Phase3TwoPart, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "Phase3TwoPartId", data.Phase3TwoPartId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *Phase3TwoPartRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "Phase3TwoPartId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
