package models

type Result struct {
	UpdateID      int            `json:"update_id"`
	Message       *Message       `json:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Entity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Message struct {
	MessageID   int         `json:"message_id"`
	From        From        `json:"from"`
	Chat        *Chat       `json:"chat"`
	Date        int         `json:"date"`
	Text        string      `json:"text"`
	Entities    []Entity    `json:"entities,omitempty"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	Location    *Location   `json:"location"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboard `json:"inline_keyboard"`
}

type InlineKeyboard struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type CallbackQuery struct {
	ID           string  `json:"id"`
	From         From    `json:"from"`
	Message      Message `json:"message"`
	ChatInstance string  `json:"chat_instance"`
	Data         string  `json:"data"`
}
