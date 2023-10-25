package group

import "testing"

type User struct {
	Id   int64
	Name string
	Age  int64
}

func Test(t *testing.T) {
	arr := []User{
		{1, "a", 100},
		{2, "b", 100},
		{3, "c", 200},
		{4, "d", 200},
	}

	groupById := GroupBy(arr, func(v User) int64 {
		return v.Id % 2
	})

}
