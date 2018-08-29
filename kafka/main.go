package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"192.168.126.5:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_log1111"
	msg.Value = sarama.StringEncoder("this is a test log from aaron")
	for {
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("Send message failed. err:", err)
			return
		}
		fmt.Printf("pid:%v, offset:%v\n", pid, offset)
		time.Sleep(time.Second)
	}

	return
}
