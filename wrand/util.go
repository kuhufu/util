package wrand

func pow(x int64, n int64) int64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func fastPow(x int64, n int64) int64 {
	if n == 0 {
		return 1
	}

	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return x * half * half
	}
}
