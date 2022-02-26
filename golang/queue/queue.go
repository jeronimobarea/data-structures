package queue

import "errors"

type Queue []interface{}

func (q *Queue) EnQueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) DeQueue() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("empty queue")
	}
	element := q.Peek()
	*q = (*q)[1:]
	return element, nil
}

func (q *Queue) Empty() bool {
	if q == nil {
		return true
	}
	return q.size() == 0
}

func (q Queue) Peek() interface{} {
	if q.size() == 0 {
		return nil
	}
	return q[0]
}

func (q Queue) size() int {
	return len(q)
}
