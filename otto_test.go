package otto

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup
var assertValue = make(map[string]string)

type MySub struct {
	id string
}

func (s *MySub) Handle(payload interface{}) {
	defer wg.Done()
	assertValue[s.id] = payload.(string)
}

func Test_Publish(t *testing.T) {
	wg.Add(2)
	RegisterTopics("teste1", "teste2")
	s1 := &MySub{id: "first"}
	s2 := &MySub{id: "second"}
	s3 := &MySub{id: "third"}
	RegisterSubscribers("teste1", s1, s2)
	RegisterSubscribers("teste2", s3)
	Publish("teste1", "payload1")
	wg.Wait()
	if assertValue["first"] != "payload1" || assertValue["second"] != "payload1" {
		t.Error("Failed to read from publisher")
	}
}
