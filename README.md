# Otto
A very simple in-memory pub/sub system for Golang.

# How to use
```go

package main

import (
    "fmt"

    "github.com/maycowa/otto"
)

// You must define one (or as many as you need) 
// subscribers to sign in on a topic. A Subscriber
// is a struct that implements otto.Subscriber interface,
// implementing the method Handle(payload interface{}) 
type MySubscriber struct{}

// Any value may be passed to Handle() method
func (m *MySubscriber)Handle(payload interface{}) {
    fmt.PrintLn(payload.(string))
}


func main() {
    // First of all, you must register the topics that can be published:
    otto.RegisterTopics("topic1", "topic2") // You may register as much topics as you need

    // Then, you must register your subscribers:
    s1 := &MySubscriber{}
    s2 := &MySubscriber{}
    s3 := &MySubscriber{}
    otto.RegisterSubscribers("topic1", s1, s2, s3)

    // Finally, you may publish to a topic:
    otto.Publish("topic1", "my message")
}
```

# Using and Contributing

You may use this library as you wish and, if you find any bug or have any suggestion, please feel free to create a pull request =)
