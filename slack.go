package slack_wh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type config struct {
	URL string
}

type Slack interface {
	SendMessage(message Message) error
}

type slack struct {
	config *config
}

func New(url string) Slack {
	return &slack{
		config: &config{
			URL: url,
		},
	}
}

func (s *slack) SendMessage(message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	resp, err := http.Post(s.config.URL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error status code: %v", resp.StatusCode)
	}
	return nil
}
