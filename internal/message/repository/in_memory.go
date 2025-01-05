package message_repository

type InMemory struct {
	messages []string
}

func NewInMemory() *InMemory {
	return &InMemory{
		messages: make([]string, 0),
	}
}

func (m *InMemory) Create(message string) {
	m.messages = append(m.messages, message)
}

func (m *InMemory) Messages() []string {
	return m.messages
}
