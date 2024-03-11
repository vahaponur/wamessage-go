package wamessage

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type WaMessage struct {
	RegistrationID string `json:"registration-id"`
	ApiKey         string `json:"api-key"`
}

func (w WaMessage) SendSingle(message string, number string) (*resty.Response, error) {
	client := resty.New()
	contactsJSON := fmt.Sprintf(`[{"message":"%s","phone":"%s"}]`, message, number)
	resp, err := client.R().
		SetHeader("registration-id", w.RegistrationID).
		SetHeader("api-key", w.ApiKey).
		SetMultipartFormData(map[string]string{
			"contacts": contactsJSON,
			"type":     "1",
		}).
		Post("https://api.wamessage.app/api/v2/send/multi")
	fmt.Println(client.Header.Get("Content-Type"))
	return resp, err
}
