package magic

import "container/list"

type List interface {
	PushBack(value any) *list.Element
	Remove(element *list.Element) any
}

func Test(values List) {
	first := values.PushBack(10)
	values.PushBack(20)
	last := values.PushBack(30)
	values.Remove(first)
	values.Remove(last)
}
