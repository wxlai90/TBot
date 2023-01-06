package tbot

var updateOffset = 0

type Update struct {
	Text   string
	ChatID int
	Type   string
}

func (u *Update) ReplyTextMessage(text string) {
	sendMessage(u.ChatID, text)
}

func getTbotUpdates() []Update {
	updates, err := getUpdates(updateOffset)
	if err != nil {
		// log.Printf("Error getting updates, error=%s\n", err)
		panic(err)
	}

	tbotUpdates := []Update{}
	for _, update := range updates.Result {
		tbotUpdates = append(tbotUpdates, Update{Text: update.Message.Text, ChatID: update.Message.From.ID, Type: TEXT})
		updateOffset = update.UpdateID + 1
	}

	return tbotUpdates
}
