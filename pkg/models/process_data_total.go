package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProcessDataTotal struct {
	Id                    *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	DomainId              string              `json:"domain_id,omitempty" bson:"domain_id,omitempty"`
	PositionItem          int32               `json:"position_item,omitempty" bson:"position_item,omitempty"`
	ProcessDataTotalId    string              `json:"process_data_total_id,omitempty" bson:"process_data_total_id,omitempty"`
	ProcessedData         int32               `json:"processed_data,omitempty" bson:"processed_data,omitempty"`
	PositionItemOriginal  int32               `json:"position_item_original,omitempty" bson:"position_item_original,omitempty"`
	PositionItemOriginal1 int32               `json:"position_item_original_1,omitempty" bson:"position_item_original_1,omitempty"`
	PositionItemOriginal2 int32               `json:"position_item_original_2,omitempty" bson:"position_item_original_2,omitempty"`
	CreatedTime           int32               `json:"created_time" bson:"created_time,omitempty"`
	UpdatedTime           int32               `json:"updated_time" bson:"updated_time,omitempty"`
}
