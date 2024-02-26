package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	Client.HSet(Ctx, "user-1", "id", "1")
	Client.HSet(Ctx, "user-1", "name", "farid anang s")
	Client.HSet(Ctx, "user-1", "email", "faridanangs@gmail.com")

	user := Client.HGetAll(Ctx, "user-1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "farid anang s", user["name"])
	assert.Equal(t, "faridanangs@gmail.com", user["email"])

	Client.HDel(Ctx, "user-1")
}
