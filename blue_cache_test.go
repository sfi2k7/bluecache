package bluecache

import (
	"testing"
	"time"
)

func TestCacheSet(t *testing.T) {
	c := New()
	err := c.Set([]byte("Key"), []byte("Value"))
	if err != nil {
		t.Error("Errror setting value")
	}
}

func TestCacheSetEx(t *testing.T) {
	c := New()
	err := c.SetEx([]byte("Key"), []byte("Value"), time.Second*1)
	if err != nil {
		t.Error("Errror setting value")
	}
}

func TestCacheGet(t *testing.T) {
	c := New()
	err := c.Set([]byte("Key"), []byte("Value"))
	if err != nil {
		t.Error("Errror setting value")
	}
	v, err := c.Get([]byte("Key"))
	if err != nil {
		t.Error("Erorr getting value")
	}
	if string(v) != "Value" {
		t.Error("Value Does not match")
	}
}

func TestCacheGetEx(t *testing.T) {
	c := New()
	err := c.SetEx([]byte("Key"), []byte("Value"), time.Second*1)
	if err != nil {
		t.Error("Errror setting value")
	}
	time.Sleep(time.Second * 2)
	_, err = c.Get([]byte("Key"))
	if err == nil {
		t.Error("Key did not expire")
	}
}

func TestCacheClear(t *testing.T) {
	c := New()
	c.Clear()
}

func TestCacheDel(t *testing.T) {
	c := New()
	if c == nil {
		t.Error("Could not get neew instance")
	}
	err := c.Set([]byte("Key"), []byte("Value"))
	if err != nil {
		t.Error("Errror setting value")
	}
	err = c.Del([]byte("Key"))
	if err != nil {
		t.Error("Error deleting key")
	}
}
