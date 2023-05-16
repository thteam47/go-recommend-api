package defaultcomponent

import (
	"github.com/thteam47/common-libs/confg"
	"github.com/thteam47/common-libs/mongoutil"
	"github.com/thteam47/common/entity"
	"github.com/thteam47/common/pkg/mongorepository"
	"github.com/thteam47/go-recommend-api/errutil"
	"github.com/thteam47/go-recommend-api/pkg/models"
)

type ProcessDataRtbRepository struct {
	config         *ProcessDataRtbRepositoryConfig
	baseRepository *mongorepository.BaseRepository
}

type ProcessDataRtbRepositoryConfig struct {
	MongoClientWrapper *mongoutil.MongoClientWrapper `mapstructure:"mongo-client-wrapper"`
}

func NewProcessDataRtbRepositoryWithConfig(properties confg.Confg) (*ProcessDataRtbRepository, error) {
	config := ProcessDataRtbRepositoryConfig{}
	err := properties.Unmarshal(&config)
	if err != nil {
		return nil, errutil.Wrap(err, "Unmarshal")
	}

	mongoClientWrapper, err := mongoutil.NewBaseMongoClientWrapperWithConfig(properties.Sub("mongo-client-wrapper"))
	if err != nil {
		return nil, errutil.Wrap(err, "NewBaseMongoClientWrapperWithConfig")
	}
	return NewProcessDataRtbRepository(&ProcessDataRtbRepositoryConfig{
		MongoClientWrapper: mongoClientWrapper,
	})
}

func NewProcessDataRtbRepository(config *ProcessDataRtbRepositoryConfig) (*ProcessDataRtbRepository, error) {
	inst := &ProcessDataRtbRepository{
		config: config,
	}

	var err error
	inst.baseRepository, err = mongorepository.NewBaseRepository(&mongorepository.BaseRepositoryConfig{
		MongoClientWrapper: inst.config.MongoClientWrapper,
		Prototype:          models.ProcessDataRtb{},
		MongoIdField:       "Id",
		IdField:            "ProcessDataRtbId",
	})
	if err != nil {
		return nil, errutil.Wrap(err, "mongorepository.NewBaseRepository")
	}

	return inst, nil
}
func (inst *ProcessDataRtbRepository) FindAll(userContext entity.UserContext, findRequest *entity.FindRequest) ([]models.ProcessDataRtb, error) {
	result := []models.ProcessDataRtb{}
	err := inst.baseRepository.FindAll(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindAll")
	}
	return result, nil
}

func (inst *ProcessDataRtbRepository) Count(userContext entity.UserContext, findRequest *entity.FindRequest) (int32, error) {
	result, err := inst.baseRepository.Count(userContext, findRequest, &mongorepository.FindOptions{})
	if err != nil {
		return 0, errutil.Wrap(err, "baseRepository.Count")
	}
	return int32(result), nil
}

func (inst *ProcessDataRtbRepository) FindById(userContext entity.UserContext, id string) (*models.ProcessDataRtb, error) {
	result := &models.ProcessDataRtb{}
	err := inst.baseRepository.FindOneByAttribute(userContext, "ProcessDataRtbId", id, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataRtbRepository) FindByOneByAttribute(userContext entity.UserContext, filters map[string]interface{}) (*models.ProcessDataRtb, error) {
	result := &models.ProcessDataRtb{}
	err := inst.baseRepository.FindOneByAttributes(userContext, filters, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByAttribute")
	}
	return result, nil
}

func (inst *ProcessDataRtbRepository) FindByOneByFindRequest(userContext entity.UserContext, findRequest *entity.FindRequest) (*models.ProcessDataRtb, error) {
	result := &models.ProcessDataRtb{}
	err := inst.baseRepository.FindOneByFindRequest(userContext, findRequest, &mongorepository.FindOptions{}, &result)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.FindOneByFindRequest")
	}
	return result, nil
}

func (inst *ProcessDataRtbRepository) FindByTenantId(userContext entity.UserContext, tenantId string) (*models.ProcessDataRtb, error) {
	result := &models.ProcessDataRtb{}
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
func (inst *ProcessDataRtbRepository) Create(userContext entity.UserContext, data *models.ProcessDataRtb) (*models.ProcessDataRtb, error) {
	err := inst.baseRepository.Create(userContext, data, nil)
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.Create")
	}
	return data, nil
}

func (inst *ProcessDataRtbRepository) CreateAndUpdate(userContext entity.UserContext, data *models.ProcessDataRtb) error {
	result := &models.ProcessDataRtb{}
	err := inst.baseRepository.FindOneByAttributes(userContext, map[string]interface{}{
		"DomainId":     data.DomainId,
		"PositionItemOriginal": data.PositionItemOriginal,
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
		err = inst.baseRepository.UpdateOneByAttribute(userContext, "ProcessDataRtbId", result.ProcessDataRtbId, data, nil, &mongorepository.UpdateOptions{
			ExcludedProperties: excludedProperties,
			InsertIfNotExisted: true,
		})
		if err != nil {
			return errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
		}
	}

	return nil
}

func (inst *ProcessDataRtbRepository) Update(userContext entity.UserContext, data *models.ProcessDataRtb, updateRequest *entity.UpdateRequest) (*models.ProcessDataRtb, error) {
	excludedProperties := []string{
		"CreatedTime",
	}

	err := inst.baseRepository.UpdateOneByAttribute(userContext, "ProcessDataRtbId", data.ProcessDataRtbId, data, updateRequest, &mongorepository.UpdateOptions{
		ExcludedProperties: excludedProperties,
		InsertIfNotExisted: true,
	})
	if err != nil {
		return nil, errutil.Wrap(err, "baseRepository.UpdateOneByAttribute")
	}
	return data, nil
}
func (inst *ProcessDataRtbRepository) DeleteById(userContext entity.UserContext, id string) error {
	err := inst.baseRepository.DeleteOneByAttribute(userContext, "ProcessDataRtbId", id)
	if err != nil {
		return errutil.Wrap(err, "baseRepository.DeleteOneByAttribute")
	}
	return nil
}
