package structures

type Queue[T comparable] struct {
	Vals []T
}

func NewQueue[T comparable](vals []T) *Queue[T] {
	return &Queue[T]{
		Vals: vals,
	}
}

func (q *Queue[T]) Size() int {
	return len(q.Vals)
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Push(v T) []T {
	q.Vals = append(q.Vals, v)
	return q.Vals
}

func (q *Queue[T]) Pop() (T, []T) {
	// User should check whether queue has elements or not
	if q.Size() == 0 {
		panic("Check Queue size before Pop'ing value")
	}

	el := q.Vals[0]
	switch {
	case q.Size() == 1:
		q.Vals = []T{}
	case q.Size() > 1:
		q.Vals = q.Vals[1:]

	}

	return el, q.Vals
}
