package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type temp struct {
	m map[int64]bool
	lock sync.RWMutex
}

func newTemp() *temp {
	m := temp{
		m: make(map[int64]bool),
	}

	go func() {
		tick := time.NewTicker(time.Millisecond * 5)
		select {
		case <-tick.C:
			m.lock.Lock()
			m.m[rand.Int63()] = true
			m.lock.Unlock()
		}
	}()

	return &m
}

func (t *temp) getMap() map[int64]bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.m
}

func BenchmarkM(b *testing.B) {
	m := newTemp()
	var t map[int64]bool
	for n := 0; n < b.N; n++ {
		t = m.getMap()
	}
	m.lock.RLock()
	fmt.Println(t)
	m.lock.RUnlock()
}
