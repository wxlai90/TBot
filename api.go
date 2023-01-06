package tbot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/wxlai90/tbot/models"
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

func getUpdates(offset int) (*models.Update, error) {
	req, err := http.NewRequest(http.MethodGet, UPDATES_URL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("offset", fmt.Sprint(offset))
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	update := &models.Update{}
	err = json.Unmarshal(body, update)
	if err != nil {
		return nil, err
	}

	return update, nil
}
