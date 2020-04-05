package main

import "github.com/gin-gonic/gin"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/route"

func main() {
	router := gin.Default()

	v1 := router.Group("/v1") 
	{
		v1.GET("/", route.IndexRoute)
	}

	router.Run()
}