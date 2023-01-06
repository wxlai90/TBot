package tbot

import (
	"fmt"
	"time"
)

type TBot struct {
	token    string
	handlers map[string]Handler
}

func (t *TBot) SendTextMessage(chatID int, text string) {
	sendMessage(chatID, text)
}

func (t *TBot) Listen() {
	for {
		updates := getTbotUpdates()
		for _, update := range updates {
			if handler, ok := t.handlers[update.Type]; ok {
				handler(t, update)
			}
		}
		time.Sleep(time.Second * 1)
	}
}

func New(token string) *TBot {
	initEndpoints(token)

	return &TBot{
		token:    token,
		handlers: map[string]Handler{},
	}
}

func initEndpoints(token string) {
	SEND_MESSAGE = fmt.Sprintf(SEND_MESSAGE, token)
	UPDATES_URL = fmt.Sprintf(UPDATES_URL, token)
}
