package adt

import "errors"

type LinkedList struct {
	Element interface{}
	next    *LinkedList
}

func NewLinkedList(element interface{}) *LinkedList {
	return &LinkedList{element, nil}
}

func (list *LinkedList) HasNext() bool {
	return list.next != nil
}

func (list *LinkedList) Next() (*LinkedList, error) {
	if !list.HasNext() {
		return nil, errors.New("no next element")
	}
	return list.next, nil
}

func (list *LinkedList) Len() int {
	if !list.HasNext() {
		return 1
	}
	return 1 + list.next.Len()
}

func (list *LinkedList) Insert(element interface{}, index int) (*LinkedList, error) {
	if index <= 0 {
		return list.insertNext(element), nil
	}
	if !list.HasNext() {
		return nil, errors.New("insert index must be equal or less than len")
	}
	return list.next.Insert(element, index-1)
}

func (list *LinkedList) Remove(index int) (interface{}, error) {
	// Cannot change reference of itself to nil
	if index <= 0 {
		return nil, errors.New("cannot remove first element")
	} else if index == 1 {
		return list.removeNext()
	} else {
		return list.next.Remove(index - 1)
	}
}

func (list *LinkedList) insertNext(element interface{}) *LinkedList {
	var next *LinkedList
	if list.HasNext() {
		next = list.next
	}
	list.next = &LinkedList{element, next}
	return list.next
}

func (list *LinkedList) removeNext() (interface{}, error) {
	if !list.HasNext() {
		return nil, errors.New("list index out of bound")
	}
	var next *LinkedList
	target := list.next.Element
	// Preserves the next element of remove item
	if list.next.HasNext() {
		next = list.next.next
	}
	// The next element of remove item is now the next element of this element
	list.next = next
	return target, nil
}
