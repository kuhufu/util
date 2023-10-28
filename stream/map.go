package stream

func MapValue[K comparable, V, T any](m map[K]V, fn func(V) T) map[K]T {
	ms := map[K]T{}
	for k, v := range m {
		ms[k] = fn(v)
	}

	return ms
}

func MapKey[K, K2 comparable, V any](m map[K]V, fn func(K) K2) map[K2]V {
	ms := map[K2]V{}
	for k, v := range m {
		ms[fn(k)] = v
	}

	return ms
}

func GroupByKey[K, K2 comparable, V any](m map[K]V, fn func(K) K2) map[K2][]V {
	ms := map[K2][]V{}
	for k, v := range m {
		ms[fn(k)] = append(ms[fn(k)], v)
	}

	return ms
}

func FlatValue[K comparable, V any](m map[K][][]V) map[K][]V {
	ret := make(map[K][]V, len(m))
	for k, v := range m {
		ret[k] = append(ret[k], Flat(v)...)
	}

	return ret
}

func Values[K comparable, V any](m map[K]V) []V {
	var arr []V
	for _, v := range m {
		arr = append(arr, v)
	}

	return arr
}

func Keys[K comparable, V any](m map[K]V) []K {
	var arr []K
	for k := range m {
		arr = append(arr, k)
	}

	return arr
}

func PickBy[K comparable, V any](in map[K]V, fn func(K, V) bool) map[K]V {
	var ret = make(map[K]V)

	for k, v := range in {
		if fn(k, v) {
			ret[k] = v
		}
	}

	return ret
}

func MapReverse[K comparable, V any, T comparable](m map[K]V, fn func(V) T) map[T]K {
	ms := map[T]K{}
	for k, v := range m {
		ms[fn(v)] = k
	}

	return ms
}

// m1 1:N
// m2 1:N
// left join
func Join1NAnd1N[K comparable, V comparable, T any](m1 map[K][]V, m2 map[V][]T) map[K][]T {
	ms := map[K][]T{}
	for k, v := range m1 {
		for _, v2 := range v {
			ms[k] = append(ms[k], m2[v2]...)
		}
	}

	return ms
}

type Entry[K any, V any] struct {
	Key K `json:"key"`
	Val V `json:"val"`
}

func ToEntries[K comparable, V any](m map[K]V) []Entry[K, V] {
	list := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		list = append(list, Entry[K, V]{
			Key: k,
			Val: v,
		})
	}

	return list
}
