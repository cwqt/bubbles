package UI

import (
	hub "github.com/simonfxr/pubsub"
)

type Event struct {
	Topic string
	Data  interface{}
}

type Bus struct {
	Publish   func(topic string, data interface{})
	Subscribe func(topic string, cb func(event Event)) func()
}

func CreateBus() Bus {
	bus := hub.NewBus()

	publish := func(topic string, data interface{}) {
		// Not entirely sure why this needs to be in a goroutine
		// but otherwise the entire program gets locked up
		go func() {
			event := Event{Topic: topic, Data: data}
			bus.Publish(topic, event)
		}()
	}

	subscribe := func(topic string, cb func(event Event)) func() {
		sub := bus.SubscribeAsync(topic, cb)

		return func() {
			// Provide utility for un-subscribing
			bus.Unsubscribe(sub)
		}
	}

	return Bus{
		Publish:   publish,
		Subscribe: subscribe,
	}
}
