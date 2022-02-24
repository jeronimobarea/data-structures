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
	var testData Stack
	testData.Push("Push")

	assert.Equal(t, Stack{"Push"}, testData)
}

func Test_Stack_Pop(t *testing.T) {
	testData := setup()

	deleted, err := testData.Pop()

	assert.NoError(t, err)
	assert.Equal(t, Stack{"hi", "bye"}, testData)
	assert.Equal(t, "hey", deleted)
}

func Test_Stack_PopEmpty(t *testing.T) {
	var testData Stack

	deleted, err := testData.Pop()

	assert.Error(t, err)
	assert.EqualError(t, err, "empty stack")
	assert.Equal(t, nil, deleted)
}

func Test_Stack_PopOneItem(t *testing.T) {
	var testData Stack
	testData.Push("hi")

	deleted, err := testData.Pop()

	assert.NoError(t, err)
	assert.Equal(t, "hi", deleted)
	assert.Equal(t, Stack{}, testData)
}

func Test_Stack_Empty(t *testing.T) {
	var testData Stack

	assert.Equal(t, true, testData.Empty())
}

func Test_Stack_NotEmpty(t *testing.T) {
	testData := setup()

	assert.Equal(t, false, testData.Empty())
}

func Test_Stack_EmptyStackEmpty(t *testing.T) {
	var testData Stack

	assert.Equal(t, true, testData.Empty())
}

func Test_Stack_Size(t *testing.T) {
	testData := setup()

	assert.Equal(t, 3, testData.Size())
}

func Test_Stack_SizeEmptyStack(t *testing.T) {
	var testData Stack

	assert.Equal(t, 0, testData.Size())
}

func Test_Stack_Top(t *testing.T) {
	testData := setup()

	assert.Equal(t, "hey", testData.Top())
}

func Test_Stack_TopEmpty(t *testing.T) {
	var testData Stack

	assert.Equal(t, nil, testData.Top())
}
