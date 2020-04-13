package main

import (
	"fmt"

	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/router"
)

func main() {

	fmt.Println("Running")

	router := router.SetupRouter()
	router.Run()
}

// Docs
// https://github.com/gin-gonic/gin/issues/1435
// https://godoc.org/gopkg.in/mgo.v2
// https://gist.github.com/asm-jaime/0bd3c294f4cb4f7a7775af6749de4b28
// https://github.com/gin-gonic/gin#grouping-routes

// Models
// https://github.com/madhums/go-gin-mgo-demo

// time.Time
// https://golang.org/pkg/time/#Time
