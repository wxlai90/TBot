package tbot

import (
	"fmt"
	"time"
)

func (t *TBot) Listen() {
	for {
		update := getUpdate()
		if handler, ok := t.handlers[update.Type]; ok {
			handler(t, update)
		}
		time.Sleep(time.Second * 1)
	}
}

func New(token string) *TBot {
	initEndpoints(token)

	return &TBot{
		token:    token,
		handlers: map[int]Handler{},
	}
}

func initEndpoints(token string) {
	SEND_MESSAGE = fmt.Sprintf(SEND_MESSAGE, token)
}
