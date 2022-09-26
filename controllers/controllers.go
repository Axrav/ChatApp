package controllers

import (
	"chatapp/helpers"
	"chatapp/types"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var Init = types.NewFlow()

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "online"})
}

func ChatWeb(c *websocket.Conn) {
	// fmt.Println(c)
	// fmt.Println(c.Params("id"))
	// fmt.Println(Users)
	Init.NewConnection <- &types.User{Id: c.Params("id"), Connection: c}
	for {
		// fmt.Println("a")
		message_type, msg, err := c.ReadMessage()
		newMessage := types.Message{Type: message_type, Body: string(msg)}
		Init.NewMessage <- &newMessage
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}

func CreateRoom(c *fiber.Ctx) error {
	code := helpers.RandomString()
	return c.JSON(fiber.Map{"RoomID": code})
}
