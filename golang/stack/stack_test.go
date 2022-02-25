package stack

import (
	"testing"

	"github.com/tj/assert"
)

func setup() Stack {
	var stack Stack
	stack.Push("hi")
	stack.Push("bye")
	stack.Push("hey")
	return stack
}

func Test_Stack_Push(t *testing.T) {
	var stack Stack
	stack.Push("Push")

	assert.Equal(t, Stack{"Push"}, stack)
}

func Test_Stack_Pop(t *testing.T) {
	stack := setup()

	deleted, err := stack.Pop()

	assert.NoError(t, err)
	assert.Equal(t, Stack{"hi", "bye"}, stack)
	assert.Equal(t, "hey", deleted)
}

func Test_Stack_PopEmpty(t *testing.T) {
	var stack Stack

	deleted, err := stack.Pop()

	assert.Error(t, err)
	assert.EqualError(t, err, "empty stack")
	assert.Equal(t, nil, deleted)
}

func Test_Stack_PopOneItem(t *testing.T) {
	var stack Stack
	stack.Push("hi")

	deleted, err := stack.Pop()

	assert.NoError(t, err)
	assert.Equal(t, "hi", deleted)
	assert.Equal(t, Stack{}, stack)
}

func Test_Stack_Empty(t *testing.T) {
	var stack Stack

	assert.Equal(t, true, stack.Empty())
}

func Test_Stack_NotEmpty(t *testing.T) {
	stack := setup()

	assert.Equal(t, false, stack.Empty())
}

func Test_Stack_EmptyStackEmpty(t *testing.T) {
	var stack Stack

	assert.Equal(t, true, stack.Empty())
}

func Test_Stack_Size(t *testing.T) {
	stack := setup()

	assert.Equal(t, 3, stack.Size())
}

func Test_Stack_SizeEmptyStack(t *testing.T) {
	var stack Stack

	assert.Equal(t, 0, stack.Size())
}

func Test_Stack_Top(t *testing.T) {
	stack := setup()

	assert.Equal(t, "hey", stack.Top())
}

func Test_Stack_TopEmpty(t *testing.T) {
	var stack Stack

	assert.Equal(t, nil, stack.Top())
}
