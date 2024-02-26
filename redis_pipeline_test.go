package main

import (
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// di redis, kita bisa mengirim beberapa perintah secara langsung tanpa harus menunggu balasan satu per satu dari redis
// hal ini juga bisa di lakukan di go-redis m3nggunakan client.Pipelined(calback)
// di dalam callback, kita bisa melakukan semua command yang akan di jalankan di dalam pipeline

func TestPipeline(t *testing.T) {
	_, err := Client.Pipelined(Ctx, func(p redis.Pipeliner) error {
		p.SetEx(Ctx, "name", "farid", 4*time.Second)
		p.SetEx(Ctx, "address", "indonesia", 4*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "farid", Client.Get(Ctx, "name").Val())
	assert.Equal(t, "indonesia", Client.Get(Ctx, "address").Val())
}
