package router

import "github.com/gin-gonic/gin"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/controllers"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/middleware"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"

func SetupRouter() *gin.Engine {
    mongo := database.MongoDB{}
    mongo.SetDefault()
    
    router := gin.Default()
    
    router.Use(middleware.MiddleDB(&mongo))
	
    v1 := router.Group("/api/v1")
    {
        v1.GET("/", controllers.GetQuestions)
        v1.POST("/question", controllers.InsertQuestion)
    }
    return router
}