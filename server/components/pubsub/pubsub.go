package pubsub

import "context"

type Channel string

type PubSub interface {
	Publish(ctx context.Context, channel Channel, message *Message) error
	Subscribe(ctx context.Context, channel Channel) (ch <-chan *Message, close func())
}
