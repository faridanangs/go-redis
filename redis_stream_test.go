package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// walaupun datanya belum di subscribe dia akan tetap di simpan di dalam stream redis

func TestStream(t *testing.T) {
	for i := 0; i < 20; i++ {
		err := Client.XAdd(Ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "anangs ke-" + strconv.Itoa(i),
				"address": "indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestDeleteStream(t *testing.T) {
	Client.XDel(Ctx, "members")
}

func TestCreateConsumergroup(t *testing.T) {

	Client.XGroupDestroy(Ctx, "members", "group1")

	Client.XGroupCreate(Ctx, "members", "group1", "0")
	Client.XGroupCreateConsumer(Ctx, "members", "group1", "customer-1")
	Client.XGroupCreateConsumer(Ctx, "members", "group1", "customer-2")
	Client.XGroupCreateConsumer(Ctx, "members", "group1", "customer-3")
}

func TestGetStream(t *testing.T) {
	result := Client.XReadGroup(Ctx, &redis.XReadGroupArgs{
		Group:    "group1",
		Consumer: "customer-1",
		Streams:  []string{"members", "0"}, // tanda > ini adalah tanda yang belum di baca atau trakhir di baca
		Count:    3,
		Block:    5 * time.Second, // waktu tunggu untuk mengambil data jika kosong program berhenti
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}
}
