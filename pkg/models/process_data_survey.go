package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProcessDataSurvey struct {
	Id                    *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	DomainId              string              `json:"domain_id,omitempty" bson:"domain_id,omitempty"`
	UserId                string              `json:"user_id,omitempty" bson:"user_id,omitempty"`
	PositionUser          int32               `json:"position_user,omitempty" bson:"position_user,omitempty"`
	PositionItem          int32               `json:"position_item,omitempty" bson:"position_item,omitempty"`
	ProcessDataSurveyId   string              `json:"process_data_survey_id,omitempty" bson:"process_data_survey_id,omitempty"`
	ProcessedData         int32               `json:"processed_data,omitempty" bson:"processed_data,omitempty"`
	PositionItemOriginal  int32               `json:"position_item_original,omitempty" bson:"position_item_original,omitempty"`
	PositionItemOriginal1 int32               `json:"position_item_original_1,omitempty" bson:"position_item_original_1,omitempty"`
	PositionItemOriginal2 int32               `json:"position_item_original_2,omitempty" bson:"position_item_original_2,omitempty"`
	Part                  int32               `json:"part,omitempty" bson:"part,omitempty"`
	CreatedTime           int32               `json:"created_time" bson:"created_time,omitempty"`
	UpdatedTime           int32               `json:"updated_time" bson:"updated_time,omitempty"`
	TenantId              string              `bson:"tenant_id,omitempty" json:"tenant_id,omitempty"`
}
