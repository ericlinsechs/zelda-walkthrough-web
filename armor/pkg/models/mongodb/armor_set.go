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

func (model *ArmorModel) AllSet() ([]models.ArmorSet, error) {
	// Define variables
	ctx := context.TODO()
	ma := []models.ArmorSet{}

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

func (model *ArmorModel) FindSet(id string) (*models.ArmorSet, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}

	// Find user by id
	result := new(models.ArmorSet)

	err = model.Collection.FindOne(context.TODO(), bson.M{"_id": p}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return result, nil
}

// func (model *ArmorModel) FindSetByName(Name string) (*models.ArmorSet, error) {
// 	// Find user by id
// 	result := new(models.ArmorSet)

// 	opts := options.FindOne().SetSort(bson.D{{"setname", Name}})

// 	err := model.Collection.FindOne(context.TODO(), bson.M{}, opts).Decode(result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, errors.New("ErrNoDocuments")
// 		}
// 		return nil, err
// 	}

// 	return result, nil
// }

func (model *ArmorModel) InsertSet(Armor *models.ArmorSet) (*mongo.InsertOneResult, error) {
	return model.Collection.InsertOne(context.TODO(), *Armor)
}

func (model *ArmorModel) UpdateSetByName(Name string, ItemID string) (*mongo.UpdateResult, error) {

	opts := options.Update().SetUpsert(false)
	update := bson.M{"$addToSet": bson.M{"tag": ItemID}}
	// update := bson.M{"$set", bson.D{{fmt.Sprintf("tag.%s", Part), ItemID}}}

	result, err := model.Collection.UpdateOne(context.TODO(), bson.M{"setname": Name}, update, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return result, nil
}

func (model *ArmorModel) InsertManySet(docs []interface{}) (*mongo.InsertManyResult, error) {
	opts := options.InsertMany().SetOrdered(true)
	return model.Collection.InsertMany(context.TODO(), docs, opts)
}

func (model *ArmorModel) DeleteSet(id string) (*mongo.DeleteResult, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}
	return model.Collection.DeleteOne(context.TODO(), bson.M{"_id": primitiveID})
}
