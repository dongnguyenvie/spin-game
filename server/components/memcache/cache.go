package memcache

import (
	"sync"
	"time"
)

const (
	DEFAULT_TTL = 0
)

type Caching interface {
	Write(k string, value interface{})
	Read(k string) interface{}
	WriteTTL(k string, value interface{}, ttl int)
}

type caching struct {
	store  map[string]interface{}
	locker *sync.RWMutex
}

func NewCaching() *caching {
	return &caching{store: make(map[string]interface{}), locker: new(sync.RWMutex)}
}

func (c *caching) Write(k string, value interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value
}

func (c *caching) Read(k string) interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.store[k]
}

// fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
// ttl is second
func (c *caching) WriteTTL(k string, value interface{}, ttl int) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value
	go func() {
		timer := time.NewTimer(time.Duration(ttl) * time.Second)
		<-timer.C
		c.Write(k, nil)
	}()
}
