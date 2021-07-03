package linkedlist

import "errors"

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
	list *List
}

type List struct {
	head *Node
	tail *Node
}

var ErrEmptyList = errors.New("list is empty")

func (n Node) Next() *Node {
	return n.next
}

func (n Node) Prev() *Node {
	return n.prev
}

func (n Node) First() *Node {
	return n.list.head
}

func (n Node) Last() *Node {
	return n.list.tail
}

func NewList(args ...interface{}) *List {
	newlist := List{head: nil, tail: nil}
	var last *Node
	last = nil
	for _, val := range args {
		newnode := Node{Val: val, next: nil, prev: last, list: &newlist}
		if newlist.head == nil {
			newlist.head = &newnode
		}
		if last != nil {
			last.next = &newnode
		}
		last = &newnode
	}
	newlist.tail = last
	return &newlist
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) PushBack(v interface{}) {
	newnode := Node{Val: v, prev: l.tail, list: l, next: nil}
	if l.tail != nil {
		l.tail.next = &newnode
		l.tail = &newnode
	} else {
		l.head, l.tail = &newnode, &newnode
	}
}

func (l *List) PopBack() (interface{}, error) {
	// empty list
	if l.tail == nil {
		return nil, ErrEmptyList
	}
	// one element
	if l.tail == l.head {
		temp := l.tail.Val
		l.head, l.tail = nil, nil
		return temp, nil
	}
	temp := l.tail.Val
	l.tail, l.tail.prev.next, l.tail.next, l.tail.prev = l.tail.prev, nil, nil, nil
	return temp, nil
}

func (l *List) PushFront(v interface{}) {
	newnode := Node{Val: v, prev: nil, list: l, next: l.head}
	if l.head != nil {
		l.head.prev = &newnode
		l.head = &newnode
	} else {
		l.head, l.tail = &newnode, &newnode
	}

}

func (l *List) PopFront() (interface{}, error) {
	// No values left
	if l.head == nil {
		return nil, ErrEmptyList
	}
	// Last value left gets popped
	if l.head == l.tail {
		temp := l.head.Val
		l.head, l.tail = nil, nil
		return temp, nil
	}
	// At least 2 values in list
	temp := l.head.Val
	l.head, l.head.next.prev, l.head.prev, l.head.next = l.head.next, nil, nil, nil
	return temp, nil
}

func (l *List) Reverse() {
	nextnode := l.head
	for nextnode != nil {
		currnode := nextnode
		nextnode = nextnode.next
		currnode.next, currnode.prev = currnode.prev, currnode.next
	}
	l.head, l.tail = l.tail, l.head
}
