package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Armor struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Location     string             `bson:"location,omitempty"`
	Cost         uint               `bson:"cost,omitempty"`
	Uses         string             `bson:"uses,omitempty"`
	ImageUrl     string             `bson:"imageurl,omitempty"`
	LastTimeEdit string             `bson:"lasttimeedit,omitempty"`
}
