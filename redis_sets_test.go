package main

import (
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// sets ini sama seperti list tapi datnya unik
// dan tidak berurutan

func TestSets(t *testing.T) {
	Client.SAdd(Ctx, "students", "farid")
	Client.SAdd(Ctx, "students", "farid")
	Client.SAdd(Ctx, "students", "anangs")
	Client.SAdd(Ctx, "students", "anangs")
	Client.SAdd(Ctx, "students", "anangs")
	Client.SAdd(Ctx, "students", "anangs")
	Client.SAdd(Ctx, "students", "wagas")

	assert.Equal(t, int64(3), Client.SCard(Ctx, "students").Val())
	assert.Equal(t, []string{"farid", "anangs", "wagas"}, Client.SMembers(Ctx, "students").Val())
}

// jika kiat gunakan sorted set kita bisa atur datanya supaya bisa berurutan
func TestSortedSet(t *testing.T) {
	// kita gunakan Zadd untuk membuat sortednya dan menggunakan redis.Z untuk memasukan data scorenya yang memilki 2 nilai
	// score ini urutannya
	// member ini valuenya
	Client.ZAdd(Ctx, "scores", redis.Z{Score: 5, Member: "anang"})
	Client.ZAdd(Ctx, "scores", redis.Z{Score: 10, Member: "samudra"})
	Client.ZAdd(Ctx, "scores", redis.Z{Score: 2, Member: "farid"})

	// ZRange unutk mendapatkan semua data Zscore
	assert.Equal(t, []string{"farid", "anang", "samudra"}, Client.ZRange(Ctx, "scores", 0, 2).Val())
	assert.Equal(t, "samudra", Client.ZPopMax(Ctx, "scores").Val()[0].Member)
	assert.Equal(t, "farid", Client.ZPopMin(Ctx, "scores").Val()[0].Member)
	assert.Equal(t, float64(5), Client.ZPopMax(Ctx, "scores").Val()[0].Score)

	// ZPopMax mengembalikan data dari member yang paling tinggi scorenya
	// ZPopMin mengembalikan data dari member yang paling kecil scorenya
	// dan disini kita gunakan val()[0].Member untuk mendapatkan data membernya dengan index ke 0 kalo [0].Score untuk mendapatkan data scorenya
}
