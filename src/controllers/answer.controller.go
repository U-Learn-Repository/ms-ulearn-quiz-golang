package controllers

import "gopkg.in/mgo.v2/bson"
import "github.com/gin-gonic/gin"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"

func InsertAnswer(c *gin.Context) {

	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}

	req := models.Answer{}

	err := c.Bind(&req)

	req.Id = bson.NewObjectId()

	if err != nil {
		c.JSON(400, gin.H{"message": "Incorrect data", "body": nil})
		return
	} else {
		err := mongo.InsertAnswer(&req)

		if err != nil {
			c.JSON(400, gin.H{"message": "error post to db", "body": nil})
		}
		c.JSON(200, gin.H{"message": "post data sucess", "body": req})
	}
}