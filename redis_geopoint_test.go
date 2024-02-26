package main

import (
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestGeoPoint(t *testing.T) {
	// kita gunakan Geoadd untuk menambahkan lokasi penjual ke dalam redis
	Client.GeoAdd(Ctx, "penjual", &redis.GeoLocation{
		Name:      "toko a",
		Latitude:  -8.552068872688888,
		Longitude: 116.08675210550922,
	})
	Client.GeoAdd(Ctx, "penjual", &redis.GeoLocation{
		Name:      "balengku",
		Latitude:  -8.549828,
		Longitude: 116.096821,
	})

	// kemudian kita gunakan geo dist untuk mendapatkan jrak antar penjual/ toko a dengan balengku
	val := Client.GeoDist(Ctx, "penjual", "toko a", "balengku", "km").Val()
	assert.Equal(t, 1.1352, val)

	// di geosearch ini kita gunakan latitud dan lognitud dari mal episentrum
	// kemudian kitagunakan geosearch untuk melihat apakah penjual berada dalamjangkauan 5km dari episentrum
	result := Client.GeoSearch(Ctx, "penjual", &redis.GeoSearchQuery{
		Latitude:   -8.59159711434153,
		Longitude:  116.10505981280878,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()

	assert.Equal(t, []string{"toko a", "balengku"}, result)

}
