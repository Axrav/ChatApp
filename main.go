package main

import (
	"chatapp/controllers"
	"chatapp/router"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := router.Setup()
	app.Use(logger.New())
	go controllers.Init.Background()
	//go Init.Background()
	app.Listen(":9898")

}
