package bluecache

import (
	"time"
)

//Backend provides backend interface for BlueCache
//It has a pluggable interface
type Backend interface {
	Clear() error
	Close() error
	Del(k []byte) error
	Get(k []byte) ([]byte, error)
	Init()
	Set(k, v []byte) error
	SetEx(k, v []byte, ex time.Duration) error
}
