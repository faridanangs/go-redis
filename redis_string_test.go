package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	Client.SetEx(Ctx, "name", "anang s", 3*time.Second)

	result, err := Client.Get(Ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "anang s", result)

	time.Sleep(5 * time.Second)

	_, err = Client.Get(Ctx, "name").Result()
	assert.NotNil(t, err)
}
