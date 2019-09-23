package adt

import "testing"

func TestQueuedList(t *testing.T) {
	one := 1
	two := 2
	three := 3
	four := 4
	five := 5
	queue := NewQueue()

	if queue.Len() != 0 {
		t.Errorf("expected length %d:, but get: %d", 0, queue.Len())
	}

	err := queue.Enqueue(one)
	if queue.Len() != 1 {
		t.Errorf("expected length %d:, but get: %d", 1, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	err = queue.Enqueue(two)
	if queue.Len() != 2 {
		t.Errorf("expected length %d:, but get: %d", 2, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	err = queue.Enqueue(three)
	if queue.Len() != 3 {
		t.Errorf("expected length %d:, but get: %d", 3, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	val, err := queue.Dequeue()
	if queue.Len() != 2 {
		t.Errorf("expected length %d:, but get: %d", 2, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if val != one {
		t.Errorf("expected element: %d, but get: %d", one, val)
	}

	err = queue.Enqueue(four)
	if queue.Len() != 3 {
		t.Errorf("expected length %d:, but get: %d", 3, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	val, err = queue.Dequeue()
	if queue.Len() != 2 {
		t.Errorf("expected length %d:, but get: %d", 2, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if val != two {
		t.Errorf("expected element: %d, but get: %d", two, val)
	}

	val, err = queue.Dequeue()
	if queue.Len() != 1 {
		t.Errorf("expected length %d:, but get: %d", 1, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if val != three {
		t.Errorf("expected element: %d, but get: %d", three, val)
	}

	val, err = queue.Dequeue()
	if queue.Len() != 0 {
		t.Errorf("expected length %d:, but get: %d", 0, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if val != four {
		t.Errorf("expected element: %d, but get: %d", four, val)
	}

	_, err = queue.Dequeue()
	if queue.Len() != 0 {
		t.Errorf("expected error")
	}

	err = queue.Enqueue(five)
	if queue.Len() != 1 {
		t.Errorf("expected length %d:, but get: %d", 1, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	val, err = queue.Dequeue()
	if queue.Len() != 0 {
		t.Errorf("expected length %d:, but get: %d", 0, queue.Len())
	}
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if val != five {
		t.Errorf("expected element: %d, but get: %d", five, val)
	}

	_, err = queue.Dequeue()
	if queue.Len() != 0 {
		t.Errorf("expected error")
	}

	_, err = queue.Dequeue()
	if queue.Len() != 0 {
		t.Errorf("expected error")
	}
}
