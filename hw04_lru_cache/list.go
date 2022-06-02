package hw04lrucache

func NewList() List {
	return new(list)
}

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
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.front == nil {
		l.front = &ListItem{Value: v}
		l.back = l.front
	} else {
		l.front.Prev = &ListItem{Value: v}
		l.front.Prev.Next = l.front
		l.front = l.front.Prev
	}
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.front == nil {
		l.front = &ListItem{Value: v}
		l.back = l.front
	} else {
		l.back.Next = &ListItem{Value: v}
		l.back.Next.Prev = l.back
		l.back = l.back.Next
	}
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	i.Prev = nil
	i.Next = l.front
	l.front.Prev = i
	l.front = i
}

func (l list) Len() int {
	size := 0
	for item := l.front; item != nil; item = item.Next {
		size++
	}
	return size
}
