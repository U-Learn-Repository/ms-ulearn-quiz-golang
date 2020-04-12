package controllers

import (
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func InsertQualification(c *gin.Context) {

	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "status": 400, "body": nil})
	}

	req := models.Qualification{}

	err := c.Bind(&req)

	req.Id = bson.NewObjectId()

	if err != nil {
		c.JSON(400, gin.H{"message": "Incorrect data", "status": 400, "body": nil})
		return
	} else {
		err := mongo.InsertQualification(&req)

		if err != nil {
			c.JSON(400, gin.H{"message": "error post to db", "status": 400, "body": nil})
		}
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": req})
	}
}

// UPDATE: /api/v1/qualification/:id
func UpdateQualification(c *gin.Context) {
	id := c.Params.ByName("id")

	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "status": 400, "body": nil})
	}

	req := models.Qualification{}

	err := c.Bind(&req)

	data, err := mongo.UpdateQualification(id, req)

	if err != nil {
		c.JSON(400, gin.H{"message": "can't get data from database", "body": nil})
	} else {
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": data})
	}
}

// DELETE: /api/v1/qualification/:id
func DeleteQualification(c *gin.Context) {
	id := c.Params.ByName("id")

	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "status": 400, "body": nil})
	}

	err := mongo.DeleteQualification(id)

	if err != nil {
		c.JSON(400, gin.H{"message": "can't get data from database", "body": nil})
	} else {
		c.JSON(200, gin.H{"message": "OK, was deleted successful", "status": 200})
	}
}
