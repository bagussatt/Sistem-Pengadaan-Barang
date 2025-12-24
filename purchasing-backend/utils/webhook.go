package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)
func SendWebhook(url string, payload interface{}) error {
	if url == "" {
		return nil 
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	return err
}
