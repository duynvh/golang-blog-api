package main

import (
	"context"
	"golang-blog-api/pubsub"
	"golang-blog-api/pubsub/pblocal"
	log "golang-blog-api/log"
	"time"
)

func main() {
	var localPb pubsub.Pubsub = pblocal.NewPubSub()

	var topic pubsub.Topic = "OrderCreated"

	sub1, close1 := localPb.Subscribe(context.Background(), topic)
	sub2, _ := localPb.Subscribe(context.Background(), topic)

	localPb.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPb.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func ()  {
		for {
			log.Print("Con1:", (<-sub1).Data())
			time.Sleep(time.Microsecond * 400)
		}
	}()

	go func ()  {
		for {
			log.Print("Con1:", (<-sub2).Data())
			time.Sleep(time.Microsecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	close1()

	localPb.Publish(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 2)
}
