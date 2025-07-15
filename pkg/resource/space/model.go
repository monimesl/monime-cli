package space

import (
	"slices"
	"sort"
)

type Space struct {
	Id    string
	Name  string
	Alias string
	URL   string
}

type List struct {
	Items []Space
}

func (l *List) Add(space Space) {
	if idx := l.IndexOf(space.Id); idx == -1 {
		l.Items = append(l.Items, space)
	} else {
		l.Items[idx] = space
	}
	sort.SliceStable(l.Items, func(i, j int) bool {
		return l.Items[i].Name < l.Items[j].Name
	})
}

func (l *List) IndexOf(idOrAlias string) int {
	return slices.IndexFunc(l.Items, func(item Space) bool {
		return item.Id == idOrAlias || item.Alias == idOrAlias
	})
}

func (l *List) GetSpace(idOrAlias string) (Space, bool) {
	idx := l.IndexOf(idOrAlias)
	if idx == -1 {
		return Space{}, false
	}
	return l.Items[idx], true
}
