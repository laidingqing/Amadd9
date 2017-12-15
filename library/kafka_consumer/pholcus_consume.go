package kafka_consumer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/streamrail/concurrent-map"
)

var concurrentMap cmap.ConcurrentMap

//StartConsumer start consume message.
func StartConsumer(consumer *kafka.Consumer, timestampField string) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run == true {
		select {
		case <-sigchan:
			printPersistence() /* TODO Remove Persistence From Consumer  */
			run = false

		case ev := <-consumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				consumer.Unassign()
			case *kafka.Message:
				go processMessage(e, timestampField)
			case kafka.Error:
				fmt.Println(e)
			}
		}
	}
}

func processMessage(message *kafka.Message, timestampField string) {
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(message.Value), &jsonMap); err != nil {
		panic(err)
	}
	concurrentMap.SetIfAbsent(string(message.Key[:]), jsonMap)
}

func printPersistence() {
	for tuple := range concurrentMap.IterBuffered() {
		fmt.Println(tuple)
	}
}
