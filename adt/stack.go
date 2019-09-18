package adt

import "errors"

type Stack interface {
	Push(obj interface{})
	Pop() (val interface{}, err error)
	Peek() (val interface{}, err error)
	Len() int
}

type SliceStack struct {
	elements []interface{}
}

func NewSliceStack(optionalCapacity ...int) *SliceStack {
	if len(optionalCapacity) == 0 {
		return &SliceStack{make([]interface{}, 0)}
	}
	return &SliceStack{make([]interface{}, 0, optionalCapacity[0])}
}

func (stack *SliceStack) Push(obj interface{}) {
	stack.elements = append(stack.elements, obj)
}

func (stack *SliceStack) Pop() (val interface{}, err error) {
	val, err = stack.Peek()
	if err != nil {
		err = errors.New("cannot pop from empty slice")
		return
	}
	stack.elements = stack.elements[:stack.Len()-1]
	return
}

func (stack *SliceStack) Peek() (val interface{}, err error) {
	if stack.Len() == 0 {
		err = errors.New("cannot peek from empty slice")
		return
	}
	val = stack.elements[stack.Len()-1]
	return
}

func (stack *SliceStack) Len() int {
	return len(stack.elements)
}

func (stack *SliceStack) Cap() int {
	return cap(stack.elements)
}

func (stack *SliceStack) Equals(another *SliceStack) bool {
	return stack.EqualsToSlice(another.elements)
}

func (stack *SliceStack) EqualsToSlice(another []interface{}) (eq bool) {
	eq = false
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	stackSize := stack.Len()
	anotherSize := len(another)

	if stackSize != anotherSize {
		return
	}

	for i := 0; i < stackSize; i++ {
		if stack.elements[i] != another[i] {
			return
		}
	}

	eq = true
	return
}
