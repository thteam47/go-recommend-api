package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phase3TwoPart struct {
	Id                   *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Phase3TwoPartId      string              `json:"phase3_two_part_id,omitempty" bson:"phase3_two_part_id,omitempty"`
	DomainId             string              `json:"domain_id,omitempty" bson:"domain_id,omitempty"`
	UserId               string              `json:"user_id,omitempty" bson:"user_id,omitempty"`
	PositionUser         int32               `json:"position_user,omitempty" bson:"position_user,omitempty"`
	PositionItem         int32               `json:"position_item,omitempty" bson:"position_item,omitempty"`
	ProcessedDataC1      string              `json:"processed_data_c1,omitempty" bson:"processed_data_c1,omitempty"`
	ProcessedDataC2      string              `json:"processed_data_c2,omitempty" bson:"processed_data_c2,omitempty"`
	PositionItemOriginal int32               `json:"position_item_original,omitempty" bson:"position_item_original,omitempty"`
	CreatedTime          int32               `json:"created_time" bson:"created_time,omitempty"`
	UpdatedTime          int32               `json:"updated_time" bson:"updated_time,omitempty"`
	Part                 int32               `bson:"part,omitempty" json:"part,omitempty"`
	TenantId             string              `bson:"tenant_id,omitempty" json:"tenant_id,omitempty"`
}
