package main

import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/router"

func main() {
    router := router.SetupRouter()
    router.Run()
}