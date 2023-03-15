package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArmorSet struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	SetName  string             `bson:"setname,omitempty"`
	Effect   string             `bson:"effect,omitempty"`
	SetBonus string             `bson:"setbonus,omitempty"`
	Tag      []string           `bson:"tag,omitempty"`
	// CreatedOn   string             `bson:"createdOn,omitempty"`
}
type ArmorItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SetName     string             `bson:"setname,omitempty"`
	Name        string             `bson:"name,omitempty"`
	HowToObtain string             `bson:"howtoobtain,omitempty"`
	Url         string             `bson:"url,omitempty"`
	ImageData   string             `bson:"imagedata,omitempty"`
	Upgrade     UpgradeLevel       `bson:"upgrade,omitempty"`
}

type UpgradeLevel struct {
	FirstUpgrade  UpgradeInfo `bson:"firstupgrade,omitempty"`
	SecondUpgrade UpgradeInfo `bson:"secondupgrade,omitempty"`
	ThirdUpgrade  UpgradeInfo `bson:"thirdupgrade,omitempty"`
	FinalUpgrade  UpgradeInfo `bson:"finalupgrade,omitempty"`
}
type UpgradeInfo struct {
	Bonus     string   `bson:"bonus,omitempty"`
	Materials []string `bson:"materials,omitempty"`
}

type ArmorImage struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Data []byte             `bson:"data"`
}
