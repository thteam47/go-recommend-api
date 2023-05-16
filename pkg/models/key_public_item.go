package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type KeyPublicItem struct {
	Id              *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	DomainId        string              `bson:"domain_id,omitempty" json:"domain_id,omitempty"`
	CreatedTime     int32               `json:"created_time" bson:"created_time,omitempty"`
	UpdatedTime     int32               `json:"updated_time" bson:"updated_time,omitempty"`
	KeyPublicItemId string              `bson:"key_public_item_id,omitempty" json:"key_public_item_id,omitempty"`
	UserId          string              `bson:"user_id,omitempty" json:"user_id,omitempty"`
	PositionUser    int32               `bson:"position_user,omitempty" json:"position_user,omitempty"`
	CategoryId      string              `bson:"category_id,omitempty" json:"category_id,omitempty"`
	PositionItem    int32               `bson:"position_item,omitempty" json:"position_item,omitempty"`
	Part            int32               `bson:"part,omitempty" json:"part,omitempty"`
	KeyPublic       string              `bson:"key_public,omitempty" json:"key_public,omitempty"`
	TenantId        string              `bson:"tenant_id,omitempty" json:"tenant_id,omitempty"`
}
