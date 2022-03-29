package wrand

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"testing"
)

func TestWeightRandom_Rand(t *testing.T) {
	w := New([]Item{
		{Weight: 0.21, Val: "1"},
		{Weight: 1, Val: "2"},
		{Weight: "-0.30", Val: "3"},
	})

	res := map[interface{}]int{}

	for i := 0; i < 150000; i++ {
		res[w.Rand()] += 1
	}

	fmt.Println(res)
}

func Test(t *testing.T) {
	d := decimal.RequireFromString("01111.11")
	t.Log(d.Exponent())
	t.Log(d.CoefficientInt64())
	t.Log(d.NumDigits())
	t.Log(math.Pow(10, 2))

}
