package datetime

import (
	"fmt"
	"testing"
	"time"
)

func Test_beginOfToday(t *testing.T) {
	fmt.Println(beginOfToday())
}

func Test_beginOfCurrDay(t *testing.T) {
	fmt.Println(beginOfCurrDay(time.Now()))
}

func Test_beginOfPrevDay(t *testing.T) {
	fmt.Println(beginOfPrevDay(time.Now()))
}

func Test_beginOfNextDay(t *testing.T) {
	fmt.Println(beginOfNextDay(time.Now()))
}

func Benchmark_beginOfCurrDay(b *testing.B) {
	now := time.Now()
	for i := 0; i < b.N; i++ {
		beginOfCurrDay(now)
	}
}
