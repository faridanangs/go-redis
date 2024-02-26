package main

import (
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// Kita tahu bahwa menggunakan Redis bisa melakukan transaction menggunakan perintah MULTI dan COMMIT
// Namun harus dalam koneksi yang sama
// Karena Golang Redis melakukan maintain Connection Pool secara internal, jadi kita tidak bisa dengan mudah menggunakan MULTI dan COMMIT menggunakan redis.Client
// Kita harus menggunakan function TxPipelined(), dimana di dalamnya kita bisa membuat callback function yang berisi command-command yang akan dijalankan dalam transaction

func TestTransaction(t *testing.T) {
	_, err := Client.TxPipelined(Ctx, func(p redis.Pipeliner) error {
		p.SetEx(Ctx, "name", "farid", 4*time.Second)
		p.SetEx(Ctx, "address", "indonesia", 4*time.Second)
		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "farid", Client.Get(Ctx, "name").Val())
	assert.Equal(t, "indonesia", Client.Get(Ctx, "address").Val())
}
