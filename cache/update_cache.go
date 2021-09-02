package cache

import (
	"sync"
	"sync/atomic"
	"time"
)

type UpdateFunc func() (interface{}, error)

type MyCache struct {
	val      atomic.Value
	expireAt int64
	updating int32
	expire   time.Duration

	updateFunc UpdateFunc
	initOnce   sync.Once
}

func NewUpdateCache(expire time.Duration, updateFunc UpdateFunc) *MyCache {
	if updateFunc == nil {
		panic("updateFunc cannot be nil")
	}

	c := &MyCache{
		expire:     expire,
		updateFunc: updateFunc,
	}
	return c
}

func (c *MyCache) isUpdating() bool {
	return atomic.LoadInt32(&c.updating) == 1
}

func (c *MyCache) Get() interface{} {
	c.initOnce.Do(func() {
		c.update()
	})

	if c.expireAt > time.Now().Unix() {
		return c.val.Load()
	}

	c.update()

	return c.val.Load()
}

func (c *MyCache) update() {
	if c.isUpdating() {
		return
	}

	atomic.StoreInt32(&c.updating, 1)
	defer atomic.StoreInt32(&c.updating, 0)

	v, err := c.updateFunc()
	atomic.StoreInt64(&c.expireAt, time.Now().Add(c.expire).Unix())

	if err != nil {
		return
	}
	c.val.Store(v)
}
