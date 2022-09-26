package main

import (
	"chatapp/controllers"
	"chatapp/router"
)

func main() {
	app := router.Setup()
	go controllers.Init.Background()
	//go Init.Background()
	app.Listen(":5555")

}
