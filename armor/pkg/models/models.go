package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArmorSet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	SetName   string             `bson:"setname"`
	SetEffect []string           `bson:"seteffect,omitempty"`
	Tag       *Item              `bson:"tag,omitempty"`
}
type Item struct {
	HeadGear string `bson:"headgear,omitempty"`
	BodyGear string `bson:"bodygear,omitempty"`
	LegGear  string `bson:"leggear,omitempty"`
}

type ArmorItem struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Parent         string             `bson:"parent"`
	Part           string             `bson:"part"`
	Name           string             `bson:"name"`
	Characteristic []string           `bson:"characteristic"`
	Location       string             `bson:"location"`
	Upgrade        string             `bson:"upgrade,omitempty"`
}
