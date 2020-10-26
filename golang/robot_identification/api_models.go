package main

import (
	"sync"
	"sync/atomic"
	"time"
)

const identRobots = 100

type count32 int32

func (c *count32) inc() {
	atomic.AddInt32((*int32)(c), 1)
}

func (c *count32) dec() {
	atomic.AddInt32((*int32)(c), -1)
}

func (c *count32) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

//nolint:golint,exhaustivestruct
// CreateUserStorage ...
func CreateUserStorage(timeout time.Duration) users {
	return users{m: make(map[string]*count32), timeout: timeout}
}

type users struct {
	timeout time.Duration
	m       map[string]*count32
	mx      sync.Mutex
}

func (u *users) Inc(key string) {
	u.mx.Lock()

	c, ok := u.m[key]
	if !ok {
		count := count32(0)
		u.m[key] = &count
		c = &count
	}

	c.inc()

	time.AfterFunc(u.timeout, c.dec)

	u.mx.Unlock()
}

func (u *users) CountRobots() int32 {
	u.mx.Lock()
	defer u.mx.Unlock()

	var count int32

	for _, value := range u.m {
		if value.get() > identRobots {
			count++
		}
	}

	return count
}
