package group

import (
	"golang.org/x/exp/slices"
)

type Group[T any] struct {
	Key    string
	Values []T

	NextNodes      []*Group[T]
	prevLevelNodes []*Group[T]
	curLevelNodes  []*Group[T]
}

func New[T any](values []T) *Group[T] {
	g := &Group[T]{
		Values: values,
	}
	g.curLevelNodes = []*Group[T]{g}
	return g
}

func (g *Group[T]) GroupBy(fn func(T) string) {
	var nextLevelGroups []*Group[T]

	for _, group := range g.curLevelNodes {
		nextLevelGroups = append(nextLevelGroups, by(group, fn)...)
	}
	g.prevLevelNodes = g.curLevelNodes
	g.curLevelNodes = nextLevelGroups
}

func (g *Group[T]) Sort(fn func(ig, jg *Group[T]) int) {
	if g.prevLevelNodes == nil {
		return
	}

	for _, group := range g.prevLevelNodes {
		ng := group.NextNodes
		slices.SortFunc(ng, fn)
	}
}

func by[T any](g *Group[T], fn func(u T) string) []*Group[T] {
	var m = map[string][]T{}
	for i := 0; i < len(g.Values); i++ {
		k := fn(g.Values[i])
		m[k] = append(m[k], g.Values[i])
	}

	var ret []*Group[T]
	for k, values := range m {
		ret = append(ret, &Group[T]{
			Key:    k,
			Values: values,
		})
	}
	g.NextNodes = ret

	return ret
}
