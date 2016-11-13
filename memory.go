package bluecache

import (
	"errors"
	"sync"
	"time"
)

type memoryItem struct {
	expireOn     time.Time
	IsExpireable bool
	v            []byte
}

type memory struct {
	c map[string]*memoryItem
	l sync.RWMutex
}

func (m *memory) Init() {
	go func() {
		for {
			time.Sleep(time.Second * 1)
			m.expire()
		}
	}()
}

func (m *memory) Set(k, v []byte) error {
	if m.c == nil {
		return errors.New("Cache is not initialized")
	}

	m.l.Lock()
	defer m.l.Unlock()

	m.c[string(k)] = &memoryItem{v: v}

	return nil
}

func (m *memory) SetEx(k, v []byte, ex time.Duration) error {
	if m.c == nil {
		return errors.New("Cache is not initialized")
	}

	m.l.Lock()
	defer m.l.Unlock()

	m.c[string(k)] = &memoryItem{expireOn: time.Now().Add(ex), IsExpireable: true, v: v}
	return nil
}

func (m *memory) Get(k []byte) (v []byte, err error) {
	if m.c == nil {
		return nil, errors.New("Cache is not initialized")
	}
	m.l.RLock()
	defer m.l.RUnlock()

	if val, ok := m.c[string(k)]; ok {
		return val.v, nil
	}
	return nil, errors.New("Not Found")
}

func (m *memory) Del(k []byte) error {
	if m.c == nil {
		return errors.New("Cache is not initialized")
	}
	m.l.Lock()
	defer m.l.Unlock()

	delete(m.c, string(k))
	return nil
}

func (m *memory) Clear() error {
	if m.c == nil {
		return errors.New("Cache is not initialized")
	}
	m.l.Lock()
	defer m.l.Unlock()

	m.c = make(map[string]*memoryItem)
	return nil
}

func (m *memory) expire() {
	m.l.Lock()
	defer m.l.Unlock()
	for k, v := range m.c {
		if !v.IsExpireable {
			continue
		}
		if v.expireOn.Before(time.Now()) {
			delete(m.c, k)
		}
	}
}

func (m *memory) Close() error {
	m.c = nil
	return nil
}
