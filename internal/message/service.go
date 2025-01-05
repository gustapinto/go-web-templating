package message

import (
	"errors"
	"strings"
)

type MessageRepository interface {
	Create(string)

	Messages() []string
}

type Service struct {
	repository MessageRepository
}

func NewMessage(repository MessageRepository) Service {
	return Service{
		repository: repository,
	}
}

func (m Service) Create(message string) error {
	if len(strings.TrimSpace(message)) == 0 {
		return errors.New("message cannot be empty")
	}

	m.repository.Create(message)
	return nil
}

func (m Service) GetAll() []string {
	return m.repository.Messages()
}
