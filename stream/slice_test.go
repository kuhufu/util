package stream

import (
	"github.com/kuhufu/util/pprint"
	"testing"
)

func TestFlat(t *testing.T) {
	l1 := Map([]string{"1", "2"}, func(v string) []int64 {
		return []int64{1, 2}
	})

	pprint.Println(l1)

	flat := Flat(l1)

	pprint.Println(flat)
}

func TestDedup(t *testing.T) {
	list := Dedup([]int64{1, 2, 3, 1, 1, 2})
	pprint.Println(list)
}
