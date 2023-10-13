package tbot

import (
	"github.com/wxlai90/tbot/models"
)

type TBot struct {
	token string
}

func main() {
	t := New("some token")
	for update := range t.Listen() {
		_ = update // TODO: process update
	}
}

func (t *TBot) Listen() chan *models.Update {
	return make(chan *models.Update)
}

func New(token string) *TBot {
	return &TBot{
		token: token,
	}
}
