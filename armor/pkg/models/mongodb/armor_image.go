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

func (model *ArmorModel) FindImage(id string) (*models.ArmorImage, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}

	// Find user by id
	result := new(models.ArmorImage)

	err = model.Collection.FindOne(context.TODO(), bson.M{"_id": p}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return result, nil
}

func (model *ArmorModel) FindImageByName(name string) (*models.ArmorImage, error) {
	filter := bson.M{"name": name}

	// Find user by id
	result := new(models.ArmorImage)

	err := model.Collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return result, nil
}

func (model *ArmorModel) InsertImage(ArmorImage *models.ArmorImage) (*mongo.InsertOneResult, error) {
	return model.Collection.InsertOne(context.Background(), *ArmorImage)
}

func (model *ArmorModel) UpdateImage(ArmorImage *models.ArmorImage) (*mongo.InsertOneResult, error) {
	// return model.Collection.InsertOne(context.Background(), *ArmorImage)
}
