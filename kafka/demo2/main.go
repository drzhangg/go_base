package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	client, err := sarama.NewClient(strings.Split("localhost:9092", ","), config)
	if err != nil {
		log.Fatalln("Unable to create kafka client", err)
	}

	fmt.Println(client.Brokers())
	fmt.Println(client.Topics())

	//topics, _ := client.Topics()
	//for _, topic := range topics {
	//	fmt.Println(topic)
	//}
}
