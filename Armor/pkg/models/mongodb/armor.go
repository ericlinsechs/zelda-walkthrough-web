package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ericlinsechs/go-mongodb-microservices/Armors/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArmorModel struct {
	Collection *mongo.Collection
}

func (model *ArmorModel) All() ([]models.Armor, error) {
	// Define variables
	ctx := context.TODO()
	mm := []models.Armor{}

	// Find all users
	cursor, err := model.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	return mm, err
}

func (model *ArmorModel) FindByID(id string) (*models.Armor, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}

	// Find user by id
	Armor := new(models.Armor)

	err = model.Collection.FindOne(context.TODO(), bson.M{"_id": p}).Decode(Armor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return Armor, nil
}

func (model *ArmorModel) Insert(Armor *models.Armor) (*mongo.InsertOneResult, error) {
	return model.Collection.InsertOne(context.TODO(), *Armor)
}

func (model *ArmorModel) Delete(id string) (*mongo.DeleteResult, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}
	return model.Collection.DeleteOne(context.TODO(), bson.M{"_id": primitiveID})
}
