package helpers

import (
	"chatapp/types"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/websocket/v2"
)

func NewConn(id string, con *websocket.Conn) types.User {
	return types.User{Id: id, Connection: con}
}
func NewMessage(t int, b string) types.Message {
	return types.Message{Type: t, Body: b}
}

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 7)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:7]
}
