package router

import "github.com/gin-gonic/gin"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/controllers"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
    v1 := router.Group("/api/v1")
    {
        v1.GET("/", controllers.IndexRoute)
    }
    return router
}