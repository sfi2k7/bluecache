package bluecache

import (
	"sync"
	"time"
)

const (
	BackendMemory = iota
	BackendRedis
	BackendBolt
)

type BlueCache struct {
	backend Backend
}

func (bc *BlueCache) Set(k, v []byte) error {
	return bc.backend.Set(k, v)
}

func (bc *BlueCache) SetEx(k, v []byte, ex time.Duration) error {
	return bc.backend.SetEx(k, v, ex)
}

func (bc *BlueCache) Get(k []byte) ([]byte, error) {
	return bc.backend.Get(k)
}

func (bc *BlueCache) Del(k []byte) error {
	return bc.backend.Del(k)
}

func (bc *BlueCache) Clear() error {
	return bc.backend.Clear()
}

func NewWithBackend(backendType int) *BlueCache {
	switch backendType {
	case BackendMemory:
		return New()
	case BackendRedis:
		return NewRedis()
	case BackendBolt:
		return NewBolt()
	default:
		return nil
	}
}

//New returns new instance of BlueCache
//by default it uses Memory as store
func New() *BlueCache {
	c := &BlueCache{
		backend: &memory{
			c: make(map[string]*memoryItem),
			l: sync.RWMutex{},
		},
	}
	c.backend.Init()
	return c
}

func NewRedis() *BlueCache {
	c := &BlueCache{
		backend: &redisBackend{},
	}
	c.backend.Init()
	return c
}

func NewBolt() *BlueCache {
	c := &BlueCache{
		backend: &boltBackend{},
	}
	c.backend.Init()
	return c
}
