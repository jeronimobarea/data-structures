package single

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
		Next: l.Head,
	}
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
	if int(position) == l.Size() {
		l.Append(item)
		return
	}

	node := l.Head
	for node.Next != nil {
		if position-1 == 0 {
			newNode := &Node{
				Data: item,
				Next: node.Next,
			}
			node.Next = newNode
			return
		}
		position--
		node = node.Next
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

	node := l.Head
	for node.Next != nil {
		if position-1 == 0 {
			nextNode := node.Next
			node.Next = nextNode.Next
			return nil
		}
		position--
		node = node.Next
	}
	return nil
}

func (l *LinkedList) Size() int {
	if l.empty() {
		return 0
	}

	size := 0
	node := l.Head
	for node != nil {
		size++
		node = node.Next
	}
	return size
}

func (l *LinkedList) Search(item interface{}) (*Node, bool, error) {
	if l.empty() {
		return nil, false, errors.New("attempting search on an empty list")
	}

	node := l.Head
	for node != nil {
		if node.Data == item {
			return node, true, nil
		}
		node = node.Next
	}
	return nil, false, nil
}

func (l *LinkedList) Display() string {
	if l.empty() {
		return ""
	}

	var display string
	node := l.Head
	for node != nil {
		display += fmt.Sprintf("%v -> ", node.Data)
		node = node.Next
	}
	return display
}
