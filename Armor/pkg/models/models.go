package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Armor struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Ｎame         string             `bson:"name,omitempty"`
	Location     string             `bson:"location,omitempty"`
	Cost         int                `bson:"cost,omitempty"`
	Uses         string             `bson:"uses,omitempty"`
	LastTimeEdit string             `bson:"lasttimeedit,omitempty"`
}
