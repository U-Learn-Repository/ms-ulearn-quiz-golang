package controllers

// import "fmt"
import (
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	"github.com/gin-gonic/gin"
)

// GET: /api/v1/questions
func GetQuestions(c *gin.Context) {
	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "status": 400, "body": nil})
	}

	data, err := mongo.GetQuestions()

	if err != nil {
		c.JSON(400, gin.H{"message": "can't get data from database", "status": 400, "body": nil})
	} else {
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": data})
	}
}

// GET: /api/v1/question/:id
func GetQuestionById(c *gin.Context) {
	id := c.Params.ByName("id")

	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "status": 400, "body": nil})
	}

	data, err := mongo.GetQuestionById(id)

	if err != nil {
		c.JSON(400, gin.H{"message": "can't get data from database", "body": nil})
	} else {
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": data})
	}
}

//POST: /api/v1/question
func InsertQuestion(c *gin.Context) {
	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "status": 400, "body": nil})
	}

	req := models.Question{}

	err := c.Bind(&req)

	if err != nil {
		c.JSON(400, gin.H{"message": "Incorrect data", "body": nil})
		return
	} else {
		err := mongo.InsertQuestion(&req)

		if err != nil {
			c.JSON(400, gin.H{"message": "error post to db", "body": nil})
		}
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": req})
	}
}