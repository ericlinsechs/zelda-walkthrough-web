package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArmorSet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SetName   string             `bson:"setname"`
	SetEffect string             `bson:"seteffect"`
	Items     []ArmorItem        `bson:"Items"`
}
type ArmorItem struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Characteristic []string           `bson:"characteristic"`
	Location       string             `bson:"location"`
	Upgrade        string             `bson:"upgrade,omitempty"`
}
