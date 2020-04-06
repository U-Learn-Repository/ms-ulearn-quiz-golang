package middleware

import "github.com/gin-gonic/gin"
import "gopkg.in/mgo.v2/bson"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/database"


func MiddleDB(mongo *MongoDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := mongo.SetSession()
		if err != nil {
			c.Abort()
		} else {
			c.Set("mongo", mongo)
			c.Next()
		}
	}
}