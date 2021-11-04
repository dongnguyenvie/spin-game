package local_pubsub

import (
	"context"
	"log"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/pubsub"
	"sync"
)

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Channel][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Channel][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, channel pubsub.Channel, data *pubsub.Message) error {
	data.SetChannel(channel)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Printf("New event published: %s, with data %s", data.String(), data.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, channel pubsub.Channel) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if subs, ok := ps.mapChannel[channel]; ok {
		subs := append(subs, c)
		ps.mapChannel[channel] = subs
	} else {
		ps.mapChannel[channel] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")
		if subs, ok := ps.mapChannel[channel]; ok {
			for i := range subs {
				if subs[i] == c {
					// remove element at index in chans
					subs = append(subs[:i], subs[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[channel] = subs
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		defer common.AppRecover()
		for {
			mess := <-ps.messageQueue

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
