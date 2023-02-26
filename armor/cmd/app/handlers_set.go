package main

import (
	"encoding/json"
	"net/http"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
)

func (app *application) getAllSet(c *gin.Context) {
	results, err := app.armorSet.AllSet()
	if err != nil {
		app.serverError(c, err)
	}

	for _, result := range results {
		app.infoLog.Println(result)
	}
	// Send response
	c.JSON(http.StatusOK, results)
}

func (app *application) findSet(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	result, err := app.armorSet.FindSet(id)
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

func (app *application) createSet(c *gin.Context) {
	ma := new(models.ArmorSet)

	err := json.NewDecoder(c.Request.Body).Decode(&ma)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Println(ma)

	// Insert new user
	insertResult, err := app.armorSet.InsertSet(ma)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("inserted document with ID %v", insertResult.InsertedID)
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

func (app *application) deleteSet(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	// Delete user by id
	res, err := app.armorSet.DeleteSet(id)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("deleted %v documents\n", res.DeletedCount)
}

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
