package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (model *ArmorModel) AllImage() ([]models.ArmorImage, error) {
	// Define variables
	ctx := context.TODO()
	var images []models.ArmorImage

	// Find all users
	cursor, err := model.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &images)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	return images, err
}

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

func (model *ArmorModel) UpdateImage(ArmorImage *models.ArmorImage) (*mongo.UpdateResult, error) {
	// Specify the Upsert option to insert a new document if a document matching
	// the filter isn't found.
	opts := options.Update().SetUpsert(false)
	// Define the filter and update operations.
	filter := bson.M{"_id": ArmorImage.ID}
	update := bson.M{"$set": bson.M{"name": ArmorImage.Name}}

	// Update the document.
	return model.Collection.UpdateOne(context.Background(), filter, update, opts)
}

func (model *ArmorModel) DeleteImage(id string) (*mongo.DeleteResult, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}
	return model.Collection.DeleteOne(context.TODO(), bson.M{"_id": primitiveID})
}
