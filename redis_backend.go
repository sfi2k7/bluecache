package bluecache

import (
	"time"

	"gopkg.in/redis.v3"
)

type redisBackend struct {
}

func connect() *redis.Client {
	c := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, Network: "tcp"})
	return c
}

func (m *redisBackend) Init() {

}

func (m *redisBackend) Set(k, v []byte) error {

	c := connect()
	defer c.Close()

	st := c.Set(string(k), string(v), 0)

	return st.Err()
}

func (m *redisBackend) SetEx(k, v []byte, ex time.Duration) error {

	c := connect()
	defer c.Close()

	c.Set(string(k), string(v), ex)

	return nil
}

func (m *redisBackend) Get(k []byte) (v []byte, err error) {

	c := connect()
	defer c.Close()

	cm := c.Get(string(k))

	return cm.Bytes()
}

func (m *redisBackend) Del(k []byte) error {
	c := connect()
	defer c.Close()

	ic := c.Del(string(k))

	return ic.Err()
}

func (m *redisBackend) Clear() error {
	return nil
}
