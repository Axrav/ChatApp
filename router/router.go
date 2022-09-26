package router

import (
	"chatapp/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Get("/", controllers.Home)
	app.Get("/ws/:id", websocket.New(controllers.ChatWeb))
	app.Get("/createroom", controllers.CreateRoom)
	return app
}
