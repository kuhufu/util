package wrand

import (
	"github.com/shopspring/decimal"
	"math/rand"
	"time"
)

type Item struct {
	Weight interface{}
	Val    interface{}
}

type WeightRandom struct {
	rand        *rand.Rand
	items       []Item
	weights     []int64
	totalWeight int64
	precious    int32
}

func scan(v interface{}) decimal.Decimal {
	switch t := v.(type) {
	case int:
		v = int64(t)
	case int32:
		v = int64(t)
	}

	var d decimal.Decimal
	err := d.Scan(v)
	if err != nil {
		panic(err)
	}

	return d
}

func New(items []Item) *WeightRandom {
	if len(items) == 0 {
		panic("items length can not be zero")
	}

	decimals := make([]decimal.Decimal, 0, len(items))
	for _, v := range items {
		d := scan(v.Weight)
		if !d.IsPositive() {
			panic("weight must greater than zero")
		}

		decimals = append(decimals, d)
	}

	precious := decimals[0].Exponent()
	for _, v := range decimals[1:] {
		if p := v.Exponent(); p < precious {
			precious = p
		}
	}

	if precious < 0 {
		precious = -precious
	} else {
		precious = 0
	}

	weights := make([]int64, len(items))

	var totalWeight int64
	for i := 0; i < len(items); i++ {
		weightInt := decimals[i].Mul(decimal.NewFromInt(pow(10, int64(precious)))).IntPart()
		weights[i] = weightInt
		totalWeight += weightInt
	}

	w := &WeightRandom{
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
		items:       items,
		totalWeight: totalWeight,
		weights:     weights,
		precious:    precious,
	}
	return w
}

func (w *WeightRandom) Rand() interface{} {
	n := w.rand.Int63n(w.totalWeight)

	var start int64
	for i, weight := range w.weights {
		start += weight
		if n < start {
			return w.items[i].Val
		}
	}

	return nil
}
