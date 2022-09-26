package types

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

type Room struct {
	Id string
}
type User struct {
	Id         string
	Connection *websocket.Conn
}

type Message struct {
	Type int
	Body string
}

type Flow struct {
	Users         []*User
	NewConnection chan *User
	NewMessage    chan *Message
}

func NewFlow() *Flow {
	return &Flow{
		Users:         make([]*User, 0),
		NewConnection: make(chan *User),
		NewMessage:    make(chan *Message),
	}
}

func (f *Flow) Background() {
	for {
		select {
		case newConn := <-f.NewConnection:
			f.Users = append(f.Users, newConn)
		case newMessage := <-f.NewMessage:
			for _, x := range f.Users {
				if err := x.Connection.WriteMessage(newMessage.Type, []byte(newMessage.Body)); err != nil {
					fmt.Println(err)
				}
			}

		}
	}
}
