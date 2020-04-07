package router

import (
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/controllers"
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"
	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	mongo := database.MongoDB{}
	mongo.SetDefault()

	router := gin.Default()

	router.Use(middleware.MiddleDB(&mongo))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/questions", controllers.GetQuestions)
		v1.GET("/question/:id", controllers.GetQuestionById)
		v1.POST("/question", controllers.InsertQuestion)

		v1.POST("/answer", controllers.InsertAnswer)
		v1.POST("/qualification", controllers.InsertQualification)
	}
	return router
}
