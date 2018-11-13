package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/WangJiemin/gocomm"
)

var (
	wg sync.WaitGroup
)

func main() {
	config := gocomm.ReadConfig("./app.cnf")
	kaHost := config["ka_host"]
	kaTopic := config["ka_topic"]

	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s", kaHost)}, nil)
	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions(kaTopic)
	if err != nil {
		panic(err)
	}

	for partition, _ := range partitionList {
		pc, err := consumer.ConsumePartition(kaTopic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d, offset: %d, Key: %s, Value: %s \n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)

		wg.Wait()
		consumer.Close()
	}
}
