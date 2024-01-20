package storage

import (
	"encoding/json"
	"io"
	"net/mail"
)

type Mail struct {
	From    string    `json:"from,omitempty"`
	To      string    `json:"to,omitempty"`
	Subject string    `json:"subject,omitempty"`
	Body    io.Reader `json:"-"`
}

func NewMail(headers mail.Header, body io.Reader) *Mail {
	return &Mail{
		From:    headers.Get("From"),
		To:      headers.Get("To"),
		Subject: headers.Get("Subject"),
		Body:    body,
	}
}

func (m Mail) MarshalJSON() ([]byte, error) {
	// Read the body
	bodyBytes, err := io.ReadAll(m.Body)
	if err != nil {
		return nil, err
	}

	// Marshal the mail
	type Alias Mail
	return json.Marshal(&struct {
		*Alias
		Body string `json:"body,omitempty"`
	}{
		Alias: (*Alias)(&m),
		Body:  string(bodyBytes),
	})
}
