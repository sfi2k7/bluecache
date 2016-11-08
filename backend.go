package bluecache

import (
	"time"
)

//Backend provides backend interface for BlueCache
//It has a pluggable interface
type Backend interface {
	Set(k, v []byte) error
	SetEx(k, v []byte, ex time.Duration) error
	Get(k []byte) ([]byte, error)
	Del(k []byte) error
	Init()
	Clear() error
}
