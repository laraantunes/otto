package otto

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
func RegisterSubscribers(topic string, subscribersToRegister ...Subscriber) {
	for _, subscriber := range subscribersToRegister {
		subscribers[topic] = append(subscribers[topic], subscriber)
	}
}

// Publish Allows to publish on a specific topic, throwing a broadcast for
// all subscribers of the topic
func Publish(topic string, payload interface{}) {
	if len(subscribers[topic]) == 0 {
		return
	}
	for _, subscriber := range subscribers[topic] {
		go subscriber.Handle(payload)
	}
}
