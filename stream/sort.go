package stream

func MergeCmpFunc[T any](fns ...func(T, T) int) func(a, b T) int {
	return func(a, b T) int {
		for _, fn := range fns {
			if v := fn(a, b); v != 0 {
				return v
			}
		}
		return 0
	}
}
