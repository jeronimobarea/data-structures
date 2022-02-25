package linked_list

import (
	"errors"
	"fmt"
)

type (
	Node struct {
		Data interface{}
		Next *Node
	}

	LinkedList struct {
		Head *Node
	}
)

func (l *LinkedList) empty() bool {
	return l.Head == nil
}

func (l *LinkedList) Push(item interface{}) {
	newHead := &Node{
		Data: item,
	}
	newHead.Next = l.Head
	l.Head = newHead
}

func (l *LinkedList) Append(item interface{}) {
	node := &Node{
		Data: item,
	}
	if l.empty() {
		l.Head = node
		return
	}
	list := l.Head
	for list.Next != nil {
		list = list.Next
	}
	list.Next = node
}

func (l *LinkedList) Insert(item interface{}, position uint) {
	if position == 0 {
		l.Push(item)
		return
	}

	var counter uint = 0

	node := l.Head
	for node.Next != nil {
		if counter == position-1 {
			newNode := &Node{
				Data: item,
				Next: node.Next,
			}
			node.Next = newNode
			return
		}
		node = node.Next
		counter++
	}
}

func (l *LinkedList) Pop() (*Node, error) {
	if l.empty() {
		return nil, errors.New("empty list")
	}
	element := l.Head
	l.Head = l.Head.Next
	return element, nil
}

func (l *LinkedList) Delete(position uint) error {
	if int(position) > l.Size() {
		return errors.New("position higher than list size")
	}
	if position == 0 {
		_, err := l.Pop()
		return err
	}

	var counter uint = 0

	node := l.Head
	for node.Next != nil {
		if counter == position-1 {
			node.Next = node.Next.Next
			return nil
		}
		node = node.Next
		counter++
	}
	return nil
}

func (l *LinkedList) Size() int {
	if l.empty() {
		return 0
	}

	size := 1
	node := l.Head
	for node.Next != nil {
		size++
		node = node.Next
	}
	return size
}

func (l *LinkedList) Display() string {
	if l.empty() {
		return ""
	}
	node := l.Head
	display := fmt.Sprintf("%v -> ", l.Head.Data)
	for node.Next != nil {
		node = node.Next
		display += fmt.Sprintf("%v -> ", node.Data)
	}
	return display
}
