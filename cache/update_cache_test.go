package cache

import (
	"testing"
	"time"
)

func TestNewUpdateCache(t *testing.T) {
	c := NewUpdateCache(time.Second, func() (i interface{}, err error) {
		return time.Now().Unix(), nil
	})

	for i := 0; i < 2; i++ {
		go func() {
			for {
				t.Log(c.Get())

				time.Sleep(time.Millisecond * 500)
			}
		}()
	}

	select {}
}
