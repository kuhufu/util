package queue

import "sync"

type Func func(func() error) error

func New() Func {
	var mu = sync.Mutex{}
	return func(f func() error) error {
		mu.Lock()
		defer mu.Unlock()
		return f()
	}
}
