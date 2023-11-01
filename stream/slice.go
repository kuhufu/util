package stream

import (
	"cmp"
	"slices"
)

func GroupBy[K comparable, V any](arr []V, fn func(v V) K) map[K][]V {
	var m = map[K][]V{}
	for _, e := range arr {
		k := fn(e)
		m[k] = append(m[k], e)
	}

	return m
}

func Map[T any, V any](list []T, fn func(v T) V) []V {
	ret := make([]V, 0, len(list))
	for _, v := range list {
		ret = append(ret, fn(v))
	}

	return ret
}

func Flat[T any](list [][]T) []T {
	var size int
	for _, v := range list {
		size += len(v)
	}

	ret := make([]T, 0, size)
	for _, v := range list {
		ret = append(ret, v...)
	}

	return ret
}

func Dedup[T comparable](list []T) []T {
	ret := make([]T, 0, len(list))

	for _, v := range list {
		if Contains(ret, v) {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}

func FlatMap[T any, V any](list [][]T, fn func(e T) V) []V {
	var size int
	for _, v := range list {
		size += len(v)
	}

	ret := make([]V, 0, size)
	for _, v := range list {
		for _, t := range v {
			ret = append(ret, fn(t))
		}
	}

	return ret
}

func Filter[T any](list []T, fn func(e T) bool) []T {
	var ret []T
	for _, v := range list {
		if fn(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

func FilterOne[T any](list []T, fn func(e T) bool) T {
	var t T
	for _, v := range list {
		if fn(v) {
			return v
		}
	}
	return t
}

func Each[T any](list []T, fn func(v T)) {
	for _, t := range list {
		fn(t)
	}
}

func Sum[T cmp.Ordered](list []T) T {
	var ret T
	for _, v := range list {
		ret += v
	}

	return ret
}

func Contains[T comparable](list []T, in T) bool {
	for _, v := range list {
		if v == in {
			return true
		}
	}

	return false
}

func ToMap[K comparable, T, V any](arr []T, fn func(T) (K, V)) map[K]V {
	m := make(map[K]V, len(arr))

	for _, v := range arr {
		k, v := fn(v)
		m[k] = v
	}
	return m
}

func Merge[T any](arrs ...[]T) []T {
	var size int
	for _, arr := range arrs {
		size += len(arr)
	}

	ret := make([]T, 0, size)

	for _, arr := range arrs {
		ret = append(ret, arr...)
	}

	return ret
}

func SortByLess[K comparable, T any](arr []T, fn func(a, b T) bool) {
	slices.SortFunc(arr, func(a, b T) int {
		less := fn(a, b)
		if less {
			return -1
		}
		return 1
	})
}
