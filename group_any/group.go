package group

type User struct {
	Id   int64
	Name string
	Age  int64
}

type Group struct {
	Key        string
	Value      User
	NextGroups map[string]*Group
	Values     []User
}

func GroupBy(g *Group, fn func(u User) string) map[string]*Group {
	var m = map[string][]User{}
	for i := 0; i < len(g.Values); i++ {
		k := fn(g.Values[i])
		m[k] = append(m[k], g.Values[i])
	}

	var ret = map[string]*Group{}
	for k, users := range m {
		ret[k] = &Group{
			Key:    k,
			Values: users,
		}
	}
	g.NextGroups = ret

	return ret
}
