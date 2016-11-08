package bluecache

import (
	"sync"
	"testing"
	"time"
)

func TestMemorySet(t *testing.T) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	err := m.Set([]byte("Key"), []byte("Value"))
	if err != nil {
		t.Error("Issue with cache")
	}
}

func BenchmarkMemorySet(b *testing.B) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	for x := 0; x < b.N; x++ {
		m.Set([]byte("Key"), []byte("Value"))
	}
}

func TestMemorySetEx(t *testing.T) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	err := m.SetEx([]byte("Key"), []byte("Value"), time.Second*1)
	if err != nil {
		t.Error("Issue with cache")
	}
}

func BenchmarkMemorySetEx(b *testing.B) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	for x := 0; x < b.N; x++ {
		m.SetEx([]byte("Key"), []byte("Value"), time.Second*1)
	}
}

func TestMemoryGet(t *testing.T) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	err := m.Set([]byte("Key"), []byte("Value"))
	if err != nil {
		t.Error("Issue with cache")
	}

	v, err := m.Get([]byte("Key"))
	if err != nil {
		t.Error("Error getting data")
	}
	if string(v) != "Value" {
		t.Error("Value does not match")
	}
}

func TestMemoryGetEx(t *testing.T) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	m.Init()
	err := m.SetEx([]byte("Key"), []byte("Value"), time.Second*1)
	if err != nil {
		if err.Error() == "Cache is not initialized" {
			t.Error("Cache is not initialized")
		}
		t.Error("Issue with cache")
	}
	time.Sleep(time.Second * 2)
	_, err = m.Get([]byte("Key"))
	if err == nil {
		t.Error("Value Did not expire")
	}
}

func TestMemoryDel(t *testing.T) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	m.Init()
	err := m.Set([]byte("Key"), []byte("Value"))
	if err != nil {
		t.Error("Issue with cache")
	}
	err = m.Del([]byte("Key"))
	if err != nil {
		t.Error("Error Removing value")
	}

	_, err = m.Get([]byte("Key"))
	if err == nil {
		t.Error("Value Did not expire")
	}
}

func TestMemoryClear(t *testing.T) {
	m := &memory{
		c: make(map[string]*memoryItem),
		l: sync.RWMutex{},
	}
	err := m.Set([]byte("Key"), []byte("Value"))
	if err != nil {
		t.Error("Error setting value")
	}
	m.Clear()
}

func TestMemoryNotInitialized(t *testing.T) {
	m := &memory{
		l: sync.RWMutex{},
	}
	err := m.Clear()
	if err != nil {
		t.Error(err.Error())
	}
}
