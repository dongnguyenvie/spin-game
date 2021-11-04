package local_pubsub

import (
	"context"
	"log"
	"nolan/spin-game/components/pubsub"
	"testing"
	"time"
)

func TestPubsub(t *testing.T) {
	var localPS pubsub.PubSub = NewPubSub()

	var chn pubsub.Channel = "xxxx"

	sub1, close1 := localPS.Subscribe(context.Background(), chn)
	sub2, close2 := localPS.Subscribe(context.Background(), chn)

	localPS.Publish(context.Background(), chn, pubsub.NewMessage(1))
	localPS.Publish(context.Background(), chn, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Sub 1:", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Sub 2:", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	close1()
	close2()
	//
	localPS.Publish(context.Background(), chn, pubsub.NewMessage(3))
	//
	time.Sleep(time.Second * 2)
}
