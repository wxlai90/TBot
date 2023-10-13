package tbot

var (
	BASE_URL                  = "https://api.telegram.org/bot%s"
	UPDATES_URL               = "https://api.telegram.org/bot%s/getUpdates"
	SEND_MESSAGE              = "https://api.telegram.org/bot%s/sendMessage"
	ANSWER_CALLBACK_QUERY     = "https://api.telegram.org/bot%s/answerCallbackQuery"
	EDIT_MESSAGE_TEXT         = "https://api.telegram.org/bot%s/editMessageText"
	EDIT_MESSAGE_REPLY_MARKUP = "https://api.telegram.org/bot%s/editMessageReplyMarkup"
	DELETE_MESSAGE            = "https://api.telegram.org/bot%s/deleteMessage"
	FORWARD_MESSAGE           = "https://api.telegram.org/bot%s/forwardMessage"
)
