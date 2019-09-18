package adt

import "testing"

func TestNewStack(t *testing.T) {
	expectedCap := 3
	stack := NewSliceStack(expectedCap)
	if stack.Cap() != expectedCap {
		t.Errorf("expected cap: %d, but get: %d", expectedCap, stack.Cap())
	}
	if stack.Len() != 0 {
		t.Errorf("expected len: %d, but get: %d", 0, stack.Len())
	}
}

func TestStack_Push(t *testing.T) {
	expected := []interface{}{3.14, "Hello", -1, '!'}
	stack := NewSliceStack()
	stack.Push(3.14)
	stack.Push("Hello")
	stack.Push(-1)
	stack.Push('!')

	if !stack.EqualsToSlice(expected) {
		t.Errorf("stack: %v != slice: %v", stack, expected)
	}
	if stack.Len() != len(expected) {
		t.Errorf("expected len: %d, but get: %d", len(expected), stack.Len())
	}
}

func TestStack_PushAndPop(t *testing.T) {
	expected := []interface{}{3.14, "Hi"}
	stack := NewSliceStack()
	stack.Push(3.14)
	stack.Push("Hello")
	stack.Push(-1)
	stack.Push('!')
	if val, err := stack.Pop(); err != nil || val != '!' {
		t.Errorf("expected pop: %s, but get: %s", string('!'), string(val.(rune)))
	}
	if val, err := stack.Pop(); err != nil || val != -1 {
		t.Errorf("expected pop: %d, but get: %d", -1, val)
	}
	stack.Push("test")
	if val, err := stack.Pop(); err != nil || val != "test" {
		t.Errorf("expected pop: %s, but get: %s", "test", val)
	}
	if val, err := stack.Pop(); err != nil || val != "Hello" {
		t.Errorf("expected pop: %s, but get: %s", "Hello", val)
	}
	stack.Push("Hi")
	if !stack.EqualsToSlice(expected) {
		t.Errorf("stack: %v != slice: %v", stack, expected)
	}
	if stack.Len() != len(expected) {
		t.Errorf("expected len: %d, but get: %d", len(expected), stack.Len())
	}
}

func TestStack_EmptyStack(t *testing.T) {
	stack := NewSliceStack()
	if _, err := stack.Pop(); err == nil {
		t.Errorf("expect error when popping empty stack")
	}
	stack.Push(-1)
	if val, err := stack.Pop(); err != nil || val != -1 {
		t.Errorf("expected pop: %d, but get: %d", -1, val)
	}
	if _, err := stack.Pop(); err == nil {
		t.Errorf("expect error when popping empty stack")
	}

	if eq := stack.EqualsToSlice([]interface{}{}); !eq {
		t.Errorf("expect empty stack equals to the empty slice")
	}

	if eq := stack.Equals(NewSliceStack()); !eq {
		t.Errorf("expect empty stack equals to the another empty stack")
	}
	if stack.Len() != 0 {
		t.Errorf("expected len: %d, but get: %d", 0, stack.Len())
	}
}

func TestStack_Equals(t *testing.T) {
	stack1 := NewSliceStack()
	stack1.Push(3.14)
	stack1.Push("Hello")
	stack1.Push(-1)
	stack1.Push('!')

	stack2 := NewSliceStack()
	stack2.Push(3.14)
	stack2.Push("Hello")
	stack2.Push(-1)
	stack2.Push('!')

	if eq := stack1.Equals(stack2); !eq {
		t.Errorf("expect stack: %v equals to stack %v", stack1, stack2)
	}
}

func TestStack_Equals_NotEqual(t *testing.T) {
	stack1 := NewSliceStack()
	stack1.Push(3.14)
	stack1.Push("Hell")
	stack1.Push(-1)
	stack1.Push('!')

	stack2 := NewSliceStack()
	stack2.Push(3.14)
	stack2.Push("Hello")
	stack2.Push(-1)
	stack2.Push('!')

	if eq := stack1.Equals(stack2); eq {
		t.Errorf("expect stack: %v  not equals to stack %v", stack1, stack2)
	}

	_, _ = stack2.Pop()
	if eq := stack1.Equals(stack2); eq {
		t.Errorf("expect stack: %v  not equals to stack %v", stack1, stack2)
	}
}

func TestStack_Equals_NotComparable(t *testing.T) {
	stack1 := NewSliceStack()
	stack1.Push(3.14)
	stack1.Push("Hello")
	stack1.Push(func() {})
	stack1.Push('!')

	stack2 := NewSliceStack()
	stack2.Push(3.14)
	stack2.Push("Hello")
	stack2.Push(func() {})
	stack2.Push('!')

	if eq := stack1.Equals(stack2); eq {
		t.Errorf("stack should return false when it not comparable")
	}
}
