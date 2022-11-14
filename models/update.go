package models

type Update struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}
