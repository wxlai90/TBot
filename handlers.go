package tbot

type Handler func(t *TBot, update Update)

func (t *TBot) OnTextMessage(handler Handler) {
	t.handlers[TEXT] = handler
}
