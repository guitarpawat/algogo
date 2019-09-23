package adt

import "errors"

type Queue interface {
	Enqueue(ele interface{}) error
	Dequeue() (interface{}, error)
	Len() int
}

type QueuedList struct {
	root *LinkedList
}

func NewQueue() *QueuedList {
	return &QueuedList{}
}

func (q *QueuedList) Enqueue(ele interface{}) error {
	if q.Len() == 0 {
		q.root = NewLinkedList(ele)
	} else {
		newEle := NewLinkedList(ele)
		q.root, newEle.NextEle = newEle, q.root
	}
	return nil
}

func (q *QueuedList) Dequeue() (interface{}, error) {
	if q.Len() == 0 {
		return nil, errors.New("queue is empty")
	} else if q.Len() == 1 {
		ele := q.root.Element
		q.root = nil
		return ele, nil
	}
	ele, err := q.root.Remove(q.root.Len() - 1)
	if err != nil {
		return nil, errors.New("cannot dequeue: " + err.Error())
	}
	return ele, nil
}

func (q *QueuedList) Len() int {
	if q.root == nil {
		return 0
	}
	return q.root.Len()
}
