package group

import (
	"cmp"
	"github.com/kuhufu/util/pprint"
	"strconv"
	"testing"
)

type User struct {
	Id   int64
	Name string
	Age  int64
}

func Test(t *testing.T) {
	arr := []*User{
		{1, "a", 100},
		{2, "b", 100},
		{3, "c", 200},
		{4, "d", 200},
	}

	root := New(arr)

	root.GroupBy(func(user *User) string {
		return strconv.Itoa(int(user.Id % 2))
	})

	root.Sort(func(ig, jg *Group[*User]) int {
		return cmp.Compare(ig.Key, jg.Key)
	})

	root.GroupBy(func(user *User) string {
		return user.Name
	})

	root.Sort(func(ig, jg *Group[*User]) int {
		return cmp.Compare(ig.Key, jg.Key)
	})

	pprint.Println(root)
}
