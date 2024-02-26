package main

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     "172.22.67.37:6379",
	Username: "anangs",
	Password: "rahasia",
	DB:       0,
})
var Ctx = context.Background()

func TestConnection(t *testing.T) {
	assert.NotNil(t, Client)

	err := Client.Close()
	assert.Nil(t, err)
}

func TestPing(t *testing.T) {
	result, err := Client.Ping(Ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}
