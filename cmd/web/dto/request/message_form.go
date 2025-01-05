package request

import (
	"errors"
	"net/url"
	"strings"
)

type MessageForm struct {
	Message string
}

func NewMessageForm(values url.Values) MessageForm {
	return MessageForm{
		Message: values.Get("message"),
	}
}

func (mf MessageForm) Validate() error {
	if len(strings.TrimSpace(mf.Message)) == 0 {
		return errors.New("message cannot be empty")
	}

	return nil
}
