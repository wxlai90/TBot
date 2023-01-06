package tbot

type TBot struct {
	token    string
	handlers map[int]Handler
}

func (t *TBot) SendTextMessage(chatID int, text string) {
	sendMessage(chatID, text)
}
