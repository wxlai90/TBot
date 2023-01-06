package tbot

import (
	"fmt"
	"log"
	"net/http"
)

func sendMessage(chatID int, text string) {
	req, err := http.NewRequest(http.MethodGet, SEND_MESSAGE, nil)
	if err != nil {
		log.Println(err)
		return
	}

	q := req.URL.Query()
	q.Add("chat_id", fmt.Sprint(chatID))
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
}
