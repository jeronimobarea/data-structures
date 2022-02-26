package double

import (
	"github.com/tj/assert"
	"testing"
)

func setup() LinkedList {
	var list LinkedList
	list.Append("hi")
	list.Append("hey")
	return list
}

func Test_DoubleLinkedList_Append(t *testing.T) {
	list := setup()

	list.Append("xd")

	topNode := &Node{
		Data: "hi",
	}
	secondNode := &Node{
		Data: "hey",
		Prev: topNode,
	}
	topNode.Next = secondNode
	thirdNode := &Node{
		Data: "xd",
		Prev: secondNode,
	}
	secondNode.Next = thirdNode

	expected := LinkedList{
		Head: topNode,
	}
	assert.Equal(t, expected, list)
}

func Test_DoubleLinkedList_Insert(t *testing.T) {
	list := setup()

	list.Append("yay")
	list.Insert("xd", 2)

	topNode := &Node{
		Data: "hi",
	}
	secondNode := &Node{
		Data: "hey",
		Prev: topNode,
	}
	topNode.Next = secondNode
	thirdNode := &Node{
		Data: "xd",
		Prev: secondNode,
	}
	secondNode.Next = thirdNode
	fourthNode := &Node{
		Data: "yay",
		Prev: thirdNode,
	}
	thirdNode.Next = fourthNode

	expected := LinkedList{
		Head: topNode,
	}
	assert.Equal(t, expected, list)
}

func Test_DoubleLinkedList_InsertLastPosition(t *testing.T) {
	list := setup()

	list.Insert("xd", 2)

	topNode := &Node{
		Data: "hi",
	}
	secondNode := &Node{
		Data: "hey",
		Prev: topNode,
	}
	topNode.Next = secondNode
	thirdNode := &Node{
		Data: "xd",
		Prev: secondNode,
	}
	secondNode.Next = thirdNode

	expected := LinkedList{
		Head: topNode,
	}
	assert.Equal(t, expected, list)
}

func Test_DoubleLinkedList_Push(t *testing.T) {
	list := setup()

	list.Push("xd")

	topNode := &Node{
		Data: "xd",
	}
	secondNode := &Node{
		Data: "hi",
		Prev: topNode,
	}
	topNode.Next = secondNode
	thirdNode := &Node{
		Data: "hey",
		Prev: secondNode,
	}
	secondNode.Next = thirdNode

	expected := LinkedList{
		Head: topNode,
	}
	assert.Equal(t, expected, list)
}

func Test_DoubleLinkedList_Pop(t *testing.T) {
	list := setup()

	deleted, err := list.Pop()

	expected := LinkedList{
		Head: &Node{
			Data: "hey",
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, list)
	assert.Equal(t, "hi", deleted.Data)
}

func Test_DoubleLinkList_Delete(t *testing.T) {
	list := setup()

	list.Append("xd")
	list.Append("yay")

	topNode := &Node{
		Data: "hi",
	}
	secondNode := &Node{
		Data: "hey",
		Prev: topNode,
	}
	topNode.Next = secondNode
	thirdNode := &Node{
		Data: "yay",
		Prev: secondNode,
	}
	secondNode.Next = thirdNode

	expected := LinkedList{
		Head: topNode,
	}
	assert.NoError(t, list.Delete(2))
	assert.Equal(t, expected, list)
}
