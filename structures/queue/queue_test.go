package structures

import (
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := NewQueue([]int{})
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Size() != 3 {
		t.Errorf("expected size 3, got %d", q.Size())
	}

	if q.Vals[0] != 1 || q.Vals[1] != 2 || q.Vals[2] != 3 {
		t.Errorf("expected [1, 2, 3], got %v", q.Vals)
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue([]int{1, 2, 3})

	el, _ := q.Pop()
	if el != 1 {
		t.Errorf("expected 1, got %d", el)
	}

	el, _ = q.Pop()
	if el != 2 {
		t.Errorf("expected 2, got %d", el)
	}

	el, _ = q.Pop()
	if el != 3 {
		t.Errorf("expected 3, got %d", el)
	}

	if !q.Empty() {
		t.Errorf("expected queue to be empty")
	}
}

func TestQueue_Size(t *testing.T) {
	q := NewQueue([]int{1, 2, 3})

	if q.Size() != 3 {
		t.Errorf("expected size 3, got %d", q.Size())
	}

	q.Pop()
	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	q.Pop()
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	q.Pop()
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}
}

func TestQueue_Empty(t *testing.T) {
	q := NewQueue([]int{})

	if !q.Empty() {
		t.Errorf("expected queue to be empty")
	}

	q.Push(1)
	if q.Empty() {
		t.Errorf("expected queue to not be empty")
	}

	q.Pop()
	if !q.Empty() {
		t.Errorf("expected queue to be empty")
	}
}
