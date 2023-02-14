package locker

import (
	"sync"
)

type Locker struct {
	mx sync.Mutex
	m  map[string]bool
}

func (structure *Locker) Lock(key string) {
	structure.mx.Lock()
	defer structure.mx.Unlock()
	structure.m[key] = true
}

func (structure *Locker) IsLock(key string) bool {
	structure.mx.Lock()
	defer structure.mx.Unlock()
	return structure.m[key]
}

func (structure *Locker) Unlock(key string) {
	structure.mx.Lock()
	defer structure.mx.Unlock()
	structure.m[key] = false
}

func NewLocker() *Locker {
	return &Locker{
		m: make(map[string]bool),
	}
}
