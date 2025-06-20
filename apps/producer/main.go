package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

func radom() string {

	return fmt.Sprintf("Mensagem aleatória: %d", rand.Int())
}

func main() {

	config := LoadConfig()

	fmt.Printf("Producing to topic %s on broker %s\n", config.TopicName, config.BrokerURL)

	seeds := []string{config.BrokerURL}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ClientID("warpstream-producer"),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()

	topic := config.TopicName
	ctx := context.Background()

	for {
		var wg sync.WaitGroup
		wg.Add(1)
		record := &kgo.Record{Topic: topic, Value: []byte(radom())}
		cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
			defer wg.Done()
			if err != nil {
				fmt.Printf("record had a produce error: %v\n", err)
			}
		})
		wg.Wait()
	}
}
