package sync

import "testing"

func Test(t *testing.T) {
	m := Map[int64, string]{}
	m.Store(1, "1")
	m.Store(2, "2")
	m.Store(3, "3")

	m.Range(func(k int64, v string) bool {
		t.Log(k, v)
		return true
	})
}
