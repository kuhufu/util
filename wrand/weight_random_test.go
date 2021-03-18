package wrand

import (
	"fmt"
	"testing"
)

func TestWeightRandom_Rand(t *testing.T) {
	w := New([]Item{
		{20, "1"},
		{50, "2"},
		{30, "3"},
	})

	res := map[interface{}]int{}

	for i := 0; i < 100000; i++ {
		res[w.Rand()] += 1
	}

	fmt.Println(res)
}
