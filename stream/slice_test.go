package stream

import (
	"github.com/kuhufu/util/pprint"
	"slices"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	strArr := []string{"1", "2"}
	Map(strArr, func(v string) int {
		i, _ := strconv.Atoi(v)
		return i
	})

	t.Log(strArr)
}

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

func TestMerge(t *testing.T) {
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	a3 := []int{4, 5}

	merge := Merge(a1, a2, a3)
	t.Log(merge)
}

func TestCompact(t *testing.T) {
	compact := slices.Compact([]int{1, 3, 2, 3, 3})
	t.Log(compact)
}

func TestAll(t *testing.T) {
	all := All([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v > 0
	})

	t.Log(all)
}

func TestSome(t *testing.T) {
	some := Some([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v > 4
	})

	t.Log(some)
}
