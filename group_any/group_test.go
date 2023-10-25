package group

import (
	"github.com/kuhufu/util/pprint"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	arr := []User{
		{1, "a", 100},
		{2, "b", 100},
		{3, "c", 200},
		{4, "d", 200},
	}

	root := &Group{Values: arr}

	groupById := GroupBy(root, func(v User) string {
		return strconv.Itoa(int(v.Id % 2))
	})

	for _, g := range groupById {
		GroupBy(g, func(u User) string {
			return u.Name
		})
	}

	pprint.Println(root)
}
