package controllers

// import "fmt"
import "github.com/gin-gonic/gin"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"

func IndexRoute(c *gin.Context) {
	c.JSON(200, gin.H {
		"message": "hello world",
	})
}

func GetQuestions(c *gin.Context) {
	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't reach db", "body": nil})
	}

	data, err := mongo.GetQuestions()

	if err != nil {
		c.JSON(400, gin.H{"message": "can't get data from database", "body": nil})
	} else {
		c.JSON(200, gin.H{"message": "OK", "status": 200, "body": data})
	}
}

func InsertQuestion(c *gin.Context) {
	mongo, ok := c.Keys["mongo"].(*database.MongoDB)

	if !ok {
		c.JSON(400, gin.H{"message": "can't connect to db", "body": nil})
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
		c.JSON(200, gin.H{"message": "post data sucess", "body": req})
	}
}
