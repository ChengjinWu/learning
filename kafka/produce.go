package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"ali-a-inf-kafka-test11.bj:9092"}, config)

	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "gaia_ad_ad_service_dbs_job",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	for {
		// 生产消息
		timeStr := time.Now().Format("2006-01-02 15:04:05") + "kafka data test"
		msg.Value = sarama.ByteEncoder(timeStr)

		paritition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send Message Fail")
		}

		fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
		time.Sleep(time.Second)
	}
}
