package bench

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	res := Bench().Concurrency(10).Total(100).Do(func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	})

	fmt.Printf("%s", res)

	res.countTimesSort()
}
