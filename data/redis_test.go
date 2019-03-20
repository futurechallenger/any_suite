package data

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestDial(t *testing.T) {
	c, err := redis.Dial("tcp", "127.0.0.1:7001")
	if err != nil {
		t.Error("Redis connection error")
	}
	defer c.Close()
}
