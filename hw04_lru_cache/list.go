package hw04lrucache

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
	elemCount int
	head      *ListItem
	tail      *ListItem
}

func NewList() List {
	return new(list)
}

// Количество элементов в списке.
func (l *list) Len() int {
	return l.elemCount
}

// Первый пункт списка.
func (l *list) Front() *ListItem {
	return l.head
}

//Последний пункт списка.
func (l *list) Back() *ListItem {
	return l.tail
}

// Добавляем новый элемент в начало списка.
func (l *list) PushFront(fronElement interface{}) *ListItem {
	newItem := new(ListItem)
	if l.elemCount == 0 {
		newItem.Prev, newItem.Next, newItem.Value = nil, nil, fronElement
		l.tail, l.head = newItem, newItem
		l.elemCount++
	} else {
		newItem.Value = fronElement
		newItem.Next = l.head
		newItem.Prev = nil
		l.head.Prev = newItem
		l.head = newItem
		l.elemCount++
	}
	return newItem
}

// Добавить новый элемент в конец списка.
func (l *list) PushBack(backElement interface{}) *ListItem {
	newItem := new(ListItem)
	if l.elemCount == 0 {
		newItem.Prev, newItem.Next, newItem.Value = nil, nil, backElement
		l.tail, l.head = newItem, newItem
		l.elemCount++
	} else {
		newItem.Value = backElement
		newItem.Prev = l.tail
		newItem.Next = nil
		l.tail.Next = newItem
		l.tail = newItem
		l.elemCount++
	}
	return newItem
}

// Удалить элемент списка.
func (l *list) Remove(i *ListItem) {
	switch {
	case l.elemCount == 1:
		l.head, l.tail = nil, nil
	case i.Prev == nil:
		l.head = i.Next
		l.head.Prev = nil
	case i.Next == nil:
		l.tail = i.Prev
		l.tail.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.elemCount--
}

// Переместить элемент списка в начало списка.
func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}
	if i == l.tail {
		i.Prev.Next = nil
		l.tail = i.Prev
		i.Prev = nil
		i.Next = l.head
		l.head.Prev = i
		l.head = i
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		i.Prev = nil
		i.Next = l.head
		l.head.Prev = i
		l.head = i
	}
}
