package group

import "sort"

type User struct {
	Id   int64
	Name string
	Age  int64
}

type Group[T any] struct {
	Key        string
	Value      T
	NextGroups []Entry[string, *Group[T]]
	Values     []T

	prevLevelGroups []Entry[string, *Group[T]]
	curLevelGroups  []Entry[string, *Group[T]]
}

func (g *Group[T]) GroupBy(fn func(T) string) {
	var nextLevelGroups []Entry[string, *Group[T]]

	if g.curLevelGroups == nil {
		nextLevelGroups = GroupBy(g, fn)
	} else {
		for _, group := range g.curLevelGroups {
			nextLevelGroups = append(nextLevelGroups, GroupBy(group.V, fn)...)
		}
	}
	g.prevLevelGroups = g.curLevelGroups
	g.curLevelGroups = nextLevelGroups
}

func (g *Group[T]) Sort(fn func(ig, jg *Group[T]) bool) {
	sort.Slice(g.prevLevelGroups, func(i, j int) bool {
		return fn(g.prevLevelGroups[i].V, g.prevLevelGroups[j].V)
	})

}

func GroupBy[T any](g *Group[T], fn func(u T) string) []Entry[string, *Group[T]] {
	var m = map[string][]T{}
	for i := 0; i < len(g.Values); i++ {
		k := fn(g.Values[i])
		m[k] = append(m[k], g.Values[i])
	}

	var ret []Entry[string, *Group[T]]
	for k, values := range m {
		ret = append(ret, Entry[string, *Group[T]]{
			K: k,
			V: &Group[T]{
				Values: values,
			},
		})
	}
	g.NextGroups = ret

	return ret
}

type Entry[K, V any] struct {
	K K
	V V
}

func MapToSlice[K comparable, V any](m map[K]V) []Entry[K, V] {
	arr := make([]Entry[K, V], len(m))
	for k, v := range m {
		arr = append(arr, Entry[K, V]{
			K: k,
			V: v,
		})
	}
	return arr
}
