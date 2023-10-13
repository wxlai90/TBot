package tbot

import "github.com/wxlai90/tbot/models"

var updateOffset = 0

func getTbotUpdates() *models.Update {
	update, err := getUpdates(updateOffset)
	if err != nil {
		panic(err)
	}

	updateOffset = update.Result[len(update.Result)-1].UpdateID + 1

	return update
}
