package tbot

var updated = false

type Update struct {
	Text   string
	ChatID int
	Type   int
}

func (u *Update) ReplyTextMessage(text string) {
	sendMessage(u.ChatID, text)
}

func getUpdate() Update {
	// TODO: parse and determine the type

	if !updated {
		update := Update{
			Text:   "some text",
			ChatID: 340522357,
			Type:   TEXT,
		}

		updated = true
		return update
	}

	return Update{}
}
