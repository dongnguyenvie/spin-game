package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	data     interface{}
	createAt time.Time
	channel  Channel
	id       string
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:       fmt.Sprintf("%d", now.Nanosecond()),
		data:     data,
		createAt: now,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("Message %s", m.channel)
}

func (m *Message) Data() interface{} {
	return m.data
}

func (m *Message) Channel() Channel {
	return m.channel
}

func (m *Message) SetChannel(channel Channel) {
	m.channel = channel
}
