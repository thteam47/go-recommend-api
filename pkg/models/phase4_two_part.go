package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phase4TwoPart struct {
	Id                    *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Phase4TwoPartId       string              `json:"phase3_two_part_id,omitempty" bson:"phase3_two_part_id,omitempty"`
	DomainId              string              `json:"domain_id,omitempty" bson:"domain_id,omitempty"`
	UserId                string              `json:"user_id,omitempty" bson:"user_id,omitempty"`
	PositionUser          int32               `json:"position_user,omitempty" bson:"position_user,omitempty"`
	PositionItem          int32               `json:"position_item,omitempty" bson:"position_item,omitempty"`
	ProcessedDataR1       string              `json:"processed_data_r1,omitempty" bson:"processed_data_r1,omitempty"`
	ProcessedDataR2       string              `json:"processed_data_r2,omitempty" bson:"processed_data_r2,omitempty"`
	ProcessedDataR3       string              `json:"processed_data_r3,omitempty" bson:"processed_data_r3,omitempty"`
	PositionItemOriginal  int32               `json:"position_item_original,omitempty" bson:"position_item_original,omitempty"`
	PositionItemOriginal1 int32               `json:"position_item_original_1,omitempty" bson:"position_item_original_1,omitempty"`
	PositionItemOriginal2 int32               `json:"position_item_original_2,omitempty" bson:"position_item_original_2,omitempty"`
	CreatedTime           int32               `json:"created_time" bson:"created_time,omitempty"`
	UpdatedTime           int32               `json:"updated_time" bson:"updated_time,omitempty"`
	Part                  int32               `bson:"part,omitempty" json:"part,omitempty"`
	TenantId              string              `bson:"tenant_id,omitempty" json:"tenant_id,omitempty"`
}
