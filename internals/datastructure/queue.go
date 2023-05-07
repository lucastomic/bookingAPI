package datastructure

import "errors"

type Queue[T any] struct {
	array []T
}

func (q *Queue[T]) Push(el T) {
	q.array = append(q.array, el)
}

func (q *Queue[T]) Pop() (T, error) {
	if len(q.array) == 0 {
		return *new(T), errors.New("there is no elements in queue")
	}
	response := q.array[0]
	q.array = q.array[1:]
	return response, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.array) == 0
}

func NewQueue[T any](arr []T) *Queue[T] {
	return &Queue[T]{arr}
}
