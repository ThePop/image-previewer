package cache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  l.front,
		Prev:  nil,
	}

	if l.len == 0 {
		l.front = newItem
		l.back = newItem
	} else {
		l.front.Prev = newItem
		l.front = newItem
	}

	l.len++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.back,
	}

	if l.back == nil {
		l.front = newItem
		l.back = newItem
	} else {
		l.back.Next = newItem
		l.back = newItem
	}

	l.len++

	return newItem
}

func (l *list) Remove(i *ListItem) {
	l.len--
	if l.len == 0 {
		l.front = nil
		l.back = nil
		return
	}

	switch {
	case l.front == i:
		l.front = l.front.Next
		l.front.Prev = nil
	case l.back == i:
		l.back = l.back.Prev
		l.back.Next = nil
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case l.front == i:
		return
	case l.back == i:
		l.back.Prev.Next = nil
		l.back = l.back.Prev
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	i.Prev = nil
	i.Next = l.front
	l.front.Prev = i
	l.front = i
}
