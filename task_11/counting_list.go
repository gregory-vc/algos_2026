package task11

import "container/list"

type CountingList struct {
	*list.List
	addCount int
}

func NewCountingList(original *list.List) *CountingList {
	return &CountingList{List: original}
}

func (l *CountingList) PushBack(value any) *list.Element {
	l.addCount++
	return l.List.PushBack(value)
}

func (l *CountingList) AddCount() int {
	return l.addCount
}
