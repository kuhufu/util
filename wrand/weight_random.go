package wrand

import (
	"math/rand"
	"time"
)

type Item struct {
	Weight int
	Val    interface{}
}

type WeightRandom struct {
	rand        *rand.Rand
	items       []Item
	totalWeight int
}

func New(items []Item) *WeightRandom {
	var total int
	for _, v := range items {
		if v.Weight < 0 {
			panic("weight can not be zero")
		}
		total += v.Weight
	}

	w := &WeightRandom{
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
		items:       items,
		totalWeight: total,
	}
	return w
}

func (w *WeightRandom) Rand() interface{} {
	n := w.rand.Intn(w.totalWeight)

	start := 0
	for _, item := range w.items {
		start += item.Weight
		if n < start {
			return item.Val
		}
	}

	return nil
}
