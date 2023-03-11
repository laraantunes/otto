package otto

import (
	"fmt"
	"runtime"
)

// Subscriber Defines a interface for Pub/Sub subscribers on a topic
type Subscriber interface {
	Handle(payload interface{})
}

var topics []string
var subscribers = make(map[string][]Subscriber)

// RegisterTopics Allows to register on Otto all the available topics
func RegisterTopics(topicsToRegister ...string) {
	for _, topic := range topicsToRegister {
		topics = append(topics, topic)
	}
}

// RegisterSubscribers Allows to register subscribers on topics
func RegisterSubscribers(topic string, subscribersToRegister ...Subscriber) error {
	if !topicExists(topic) {
		return fmt.Errorf("Topic %s is not registered", topic)
	}
	for _, subscriber := range subscribersToRegister {
		subscribers[topic] = append(subscribers[topic], subscriber)
	}
	return nil
}

// Publish Allows to publish on a specific topic, throwing a broadcast for
// all subscribers of the topic
func Publish(topic string, payload interface{}) error {
	if !topicExists(topic) {
		return fmt.Errorf("Topic %s is not registered", topic)
	}
	if len(subscribers[topic]) == 0 {
		return nil
	}
	for _, subscriber := range subscribers[topic] {
		runtime.Gosched()
		go subscriber.Handle(payload)
	}
	return nil
}

func topicExists(topic string) bool {
	for _, t := range topics {
		if t == topic {
			return true
		}
	}
	return false
}
