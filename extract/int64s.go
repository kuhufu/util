package extract

func Int64List(n int, f func(i int) int64) []int64 {
	var ret []int64
	for i := 0; i < n; i++ {
		ret = append(ret, f(i))
	}

	return ret
}
