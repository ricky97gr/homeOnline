package kafka

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
)

func ConsumerMsg() {
	var wg sync.WaitGroup

	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Net.SASL.Enable = true
	config.Net.SASL.User = "user1"
	config.Net.SASL.Password = "testkafka"

	consumer, err := sarama.NewConsumer([]string{"192.168.0.202:32088"}, config)
	if err != nil {
		panic(err)
	}
	parts, err := consumer.Partitions("test")
	if err != nil {
		panic(err)
	}

	for _, p := range parts {
		pc, err := consumer.ConsumePartition("test", p, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			wg.Done()
			for msg := range partitionConsumer.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	err = consumer.Close()
	if err != nil {
		return
	}
}
