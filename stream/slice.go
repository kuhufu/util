package stream

import (
	"slices"
)

type Entry[K comparable, V any] struct {
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

func GroupBy[K comparable, V any](arr []V, fn func(v V) K) map[K][]V {
	var m = map[K][]V{}
	for _, e := range arr {
		k := fn(e)
		m[k] = append(m[k], e)
	}

	return m
}

func Sort[K comparable, T any](arr []T, fn func(a, b T) bool) {
	slices.SortFunc(arr, func(a, b T) int {
		less := fn(a, b)
		if less {
			return -1
		}
		return 1
	})
}
