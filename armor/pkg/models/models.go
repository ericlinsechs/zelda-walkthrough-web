package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArmorSet struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SetName     string             `bson:"setname,omitempty"`
	Effect      string             `bson:"effect,omitempty"`
	SetBonus    string             `bson:"setbonus,omitempty"`
	HowToObtain []string           `bson:"howtoobtain,omitempty"`
	Tag         Item               `bson:"tag,omitempty"`
	// CreatedOn   string             `bson:"createdOn,omitempty"`
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
