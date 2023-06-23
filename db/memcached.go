package db

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var MC *memcache.Client

func ConnectMemcached() {
	mc := memcache.New("localhost:11211")
	MC = mc
}
