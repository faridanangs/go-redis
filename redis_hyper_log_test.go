package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// hayperlog ini di gunakan untuk melihat jumlah user yang sudah berkunjung

func TestHyperLog(t *testing.T) {
	Client.PFAdd(Ctx, "visitors", "farid", "diana", "wagas")
	Client.PFAdd(Ctx, "visitors", "raika", "farid", "diana", "wagas")
	Client.PFAdd(Ctx, "visitors", "farid", "wagas", "raika", "jodi")

	result := Client.PFCount(Ctx, "visitors").Val()
	assert.Equal(t, int64(5), result)
}
