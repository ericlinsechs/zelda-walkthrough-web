package main

import (
	"encoding/json"
	"net/http"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) getAllItem(c *gin.Context) {
	results, err := app.armorItem.AllItem()
	if err != nil {
		app.serverError(c, err)
	}

	for _, result := range results {
		app.infoLog.Println(result)
	}

	// Send response
	c.JSON(http.StatusOK, results)
}

func (app *application) findItem(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	result, err := app.armorItem.FindItem(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.clientError(c, http.StatusBadRequest)
			return
		}
		// Any other error will send an internal server error
		app.serverError(c, err)
	}

	app.infoLog.Printf("Founded document with ID %v\n", id)

	// Send response
	c.JSON(http.StatusOK, result)
}

func (app *application) createItem(c *gin.Context) {
	ma := new(models.ArmorItem)

	err := json.NewDecoder(c.Request.Body).Decode(&ma)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Println(ma)

	// Insert new user
	insertResult, err := app.armorItem.InsertItem(ma)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("inserted document with ID %v", insertResult.InsertedID)

	// result, err := app.armorSet.FindSetByName(ma.Parent)
	// if err != nil {
	// 	if err.Error() == "ErrNoDocuments" {
	// 		app.clientError(c, http.StatusBadRequest)
	// 		return
	// 	}
	// 	// Any other error will send an internal server error
	// 	app.serverError(c, err)
	// }

	// app.infoLog.Printf("updated document %v", result)
	id := insertResult.InsertedID.(primitive.ObjectID).Hex()
	app.infoLog.Println(id)

	// var updatedDocument bson.M
	result, err := app.armorSet.FindSetAndUpdate(ma, id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.clientError(c, http.StatusBadRequest)
			return
		}
		// Any other error will send an internal server error
		app.serverError(c, err)
	}
	app.infoLog.Printf("updated document %v", result)
}

// func (app *application) createMany(c *gin.Context) {
// 	var newArmors []models.Armor

// 	err := json.NewDecoder(c.Request.Body).Decode(&newArmors)
// 	if err != nil {
// 		app.serverError(c, err)
// 	}

// 	docs, err := toDoc(newArmors)
// 	if err != nil {
// 		app.serverError(c, err)
// 	}

// 	res, err := app.armor.InsertMany(docs)
// 	if err != nil {
// 		app.serverError(c, err)
// 	}

// 	app.infoLog.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
// }

// func (app *application) delete(c *gin.Context) {
// 	// Get id from incoming url
// 	id := c.Param("id")

// 	// Delete user by id
// 	deleteResult, err := app.armor.Delete(id)
// 	if err != nil {
// 		app.serverError(c, err)
// 	}

// 	app.infoLog.Printf("%d armor(s) have been eliminated", deleteResult.DeletedCount)
// }

// func toDoc(newArmor []models.Armor) (docs []interface{}, err error) {

// 	for _, a := range newArmor {
// 		data, err := bson.Marshal(a)
// 		if err != nil {
// 			return docs, err
// 		}
// 		var doc *bson.D
// 		if err = bson.Unmarshal(data, &doc); err != nil {
// 			return docs, err
// 		}
// 		docs = append(docs, *doc)
// 	}
// 	// fmt.Println(docs)
// 	return docs, nil
// }

// func SlicetoDoc(newArmor []models.Armor) (docs []interface{}, err error) {

// 	for _, a := range newArmor {
// 		data, err := bson.Marshal(a)
// 		if err != nil {
// 			return docs, err
// 		}
// 		var doc *bson.D
// 		if err = bson.Unmarshal(data, &doc); err != nil {
// 			return docs, err
// 		}
// 		docs = append(docs, *doc)
// 	}
// 	// fmt.Println(docs)
// 	return docs, nil
// }
