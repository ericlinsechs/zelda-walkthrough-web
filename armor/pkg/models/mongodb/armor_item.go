package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArmorModel struct {
	Collection *mongo.Collection
}

func (model *ArmorModel) AllItem() ([]models.ArmorItem, error) {
	// Define variables
	ctx := context.TODO()
	ma := []models.ArmorItem{}

	// Find all users
	cursor, err := model.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &ma)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	return ma, err
}

func (model *ArmorModel) FindItem(id string) (*models.ArmorItem, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}

	// Find user by id
	result := new(models.ArmorItem)

	err = model.Collection.FindOne(context.TODO(), bson.M{"_id": p}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return result, nil
}

func (model *ArmorModel) InsertItem(Armor *models.ArmorItem) (*mongo.InsertOneResult, error) {
	return model.Collection.InsertOne(context.TODO(), *Armor)
}

// func (model *ArmorModel) InsertMany(docs []interface{}) (*mongo.InsertManyResult, error) {
// 	opts := options.InsertMany().SetOrdered(true)
// 	return model.Collection.InsertMany(context.TODO(), docs, opts)
// }

// func (model *ArmorModel) Delete(id string) (*mongo.DeleteResult, error) {
// 	primitiveID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
// 	}
// 	return model.Collection.DeleteOne(context.TODO(), bson.M{"_id": primitiveID})
// }
