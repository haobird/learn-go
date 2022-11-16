package messaging

import "time"

// defined msg bridge form
type Element struct {
	ClientID  string    `json:"clientid"`
	Topic     string    `json:"topic"`
	Payload   []byte    `json:"payload"`
	Timestamp time.Time `json:"ts"`
}

type Publisher interface {
	Publish(string, interface{}) error
}

type MessageHandler func(interface{}) error

type Subscriber interface {
	Subscribe(string, interface{}) error
	Unsubscribe(string) error
}

type PubSub interface {
	Publisher
	Subscriber
}
