package queue

import (
	"testing"

	"github.com/tj/assert"
)

func setup() Queue {
	var queue Queue
	queue.EnQueue("hi")
	queue.EnQueue("bye")
	queue.EnQueue("hey")
	return queue
}

func Test_Queue_EnQueue(t *testing.T) {
	var queue Queue

	queue.EnQueue("Enqueue")

	assert.Equal(t, Queue{"Enqueue"}, queue)
}

func Test_Queue_DeQueue(t *testing.T) {
	queue := setup()

	deleted, err := queue.DeQueue()

	assert.NoError(t, err)
	assert.Equal(t, Queue{"bye", "hey"}, queue)
	assert.Equal(t, "hi", deleted)
}

func Test_Queue_DeQueueEmpty(t *testing.T) {
	var queue Queue

	deleted, err := queue.DeQueue()

	assert.Error(t, err)
	assert.EqualError(t, err, "empty queue")
	assert.Equal(t, nil, deleted)
}

func Test_Queue_Peek(t *testing.T) {
	queue := setup()

	assert.Equal(t, "hi", queue.Peek())
}

func Test_Queue_PeekEmpty(t *testing.T) {
	var queue Queue

	assert.Equal(t, nil, queue.Peek())
}
