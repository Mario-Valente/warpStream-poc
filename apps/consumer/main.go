package main

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {

	Config := LoadConfig()

	seeds := []string{Config.BrokerURL}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ClientID("warpstream-consumer"),
		kgo.ConsumeTopics(Config.TopicName),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()
	ctx := context.Background()

	for {
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			panic(fmt.Sprint(errs))
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}
	}
}
