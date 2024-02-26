package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubscribePubSub(t *testing.T) {
	pubSub := Client.Subscribe(Ctx, "channel-1")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer pubSub.Close()
		defer wg.Done()
		for {
			// message, err := pubSub.ReceiveMessage(Ctx) // recieve message ini di gunakan utntuk mengambil data dan dia tidak akan pernah berhenti walaupun datanya tidak ada
			message, err := pubSub.ReceiveTimeout(Ctx, 8*time.Second) // namun jika kikta gunakan ini maka dia akan berhenti jika selama timeoutnya tidak ada data dia akan berhenti
			if err != nil {
				break
			}
			assert.Nil(t, err)
			fmt.Println(message)
		}
	}()

	wg.Wait()

}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10000; i++ {
		err := Client.Publish(Ctx, "channel-1", "pesan ke-"+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}
