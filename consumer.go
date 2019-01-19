package main

import (
	"log"
	"sync"

	nsq "github.com/nsqio/go-nsq"
)

const (
	topicName = "test_topic"
	chName    = "test_channel"
	nsqdAddr  = "127.0.0.1:4150"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topicName, chName, config)
	if err != nil {
		log.Fatal("new consumer error", err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Println("msg", string(msg.Body))

		return nil
	}))

	err = consumer.ConnectToNSQD(nsqdAddr)
	if err != nil {
		log.Fatal("connect to nsqd error", err)
	}

	log.Println("waiting for msgs..")

	wg.Wait()
}
