package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// untuk membuat list di redis kita bisa menggunakan Rpush
	// list ini di gunakan untuk menggolongkan data

	Client.RPush(Ctx, "names", "Farid")
	Client.RPush(Ctx, "names", "Anang")
	Client.RPush(Ctx, "names", "Samudra")

	// kemudian kita gunakan Lpop / Rpop untuk mengambil datanya
	// Lpop mengambil data dari sebeliah kiri dan Rpop dari kanan
	assert.Equal(t, "Farid", Client.LPop(Ctx, "names").Val())
	assert.Equal(t, "Anang", Client.LPop(Ctx, "names").Val())
	assert.Equal(t, "Samudra", Client.LPop(Ctx, "names").Val())

	Client.Del(Ctx, "names")
}
