package linked_list

import (
	"testing"

	"github.com/tj/assert"
)

func setup() LinkedList {
	var list LinkedList
	list.Append("bye")
	list.Append("xd")
	return list
}

func Test_LinkedList_Append(t *testing.T) {
	list := setup()

	list.Append("hi")

	expected := LinkedList{
		Head: &Node{
			Data: "bye",
			Next: &Node{
				Data: "xd",
				Next: &Node{
					Data: "hi",
				},
			},
		},
	}
	assert.Equal(t, expected, list)
}

func Test_LinkedList_Insert(t *testing.T) {
	list := setup()

	list.Insert("hi", 1)

	expected := LinkedList{
		Head: &Node{
			Data: "bye",
			Next: &Node{
				Data: "hi",
				Next: &Node{
					Data: "xd",
				},
			},
		},
	}
	assert.Equal(t, expected, list)
}

func Test_LinkedList_InsertZeroIndex(t *testing.T) {
	list := setup()

	list.Insert("yay", 0)

	expected := LinkedList{
		Head: &Node{
			Data: "yay",
			Next: &Node{
				Data: "bye",
				Next: &Node{
					Data: "xd",
				},
			},
		},
	}
	assert.Equal(t, expected, list)
}

func Test_LinkedList_InsertIndexOverflow(t *testing.T) {
	list := setup()

	list.Insert("yay", 999)

	expected := LinkedList{
		Head: &Node{
			Data: "bye",
			Next: &Node{
				Data: "xd",
			},
		},
	}
	assert.Equal(t, expected, list)
}

func Test_LinkedList_Push(t *testing.T) {
	list := setup()

	list.Push("yay")

	expected := LinkedList{
		Head: &Node{
			Data: "yay",
			Next: &Node{
				Data: "bye",
				Next: &Node{
					Data: "xd",
				},
			},
		},
	}
	assert.Equal(t, expected, list)
}

func Test_LinkedList_Pop(t *testing.T) {
	list := setup()

	deleted, err := list.Pop()

	expected := LinkedList{
		Head: &Node{
			Data: "xd",
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, list)
	assert.Equal(t, "bye", deleted.Data)
}

func Test_LinkedList_PopEmpty(t *testing.T) {
	var list LinkedList

	deleted, err := list.Pop()

	assert.Error(t, err)
	assert.EqualError(t, err, "empty list")
	assert.Nil(t, deleted)
}

func Test_LinkedList_Delete(t *testing.T) {
	list := setup()

	err := list.Delete(1)

	assert.NoError(t, err)

	expected := LinkedList{
		Head: &Node{
			Data: "bye",
		},
	}
	assert.Equal(t, expected, list)
}

func Test_LinkedList_DeleteIndexOverflow(t *testing.T) {
	list := setup()

	err := list.Delete(999)

	assert.Error(t, err)
	assert.EqualError(t, err, "position higher than list size")
	assert.Equal(t, setup(), list)
}

func Test_LinkedList_DeleteLastItem(t *testing.T) {
	var list LinkedList

	list.Push("xd")
	err := list.Delete(0)

	assert.NoError(t, err)
	assert.Equal(t, LinkedList{}, list)
}

func Test_LinkedList_Size(t *testing.T) {
	list := setup()

	assert.Equal(t, 2, list.Size())
}

func Test_LinkedList_SizeEmpty(t *testing.T) {
	var list LinkedList

	assert.Equal(t, 0, list.Size())
}

func Test_LinkedList_Display(t *testing.T) {
	list := setup()

	assert.Equal(t, "bye -> xd -> ", list.Display())
}

func Test_LinkedList_DisplayEmpty(t *testing.T) {
	var list LinkedList

	assert.Equal(t, "", list.Display())
}
