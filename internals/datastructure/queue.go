package datastructure

import "errors"

type Queue[T any] struct {
	array []T
}

// This method adds an element to the end of the queue.
func (q *Queue[T]) Push(el T) {
	q.array = append(q.array, el)
}

// This method removes the first element from the queue and returns it.
// If the queue is empty, it returns an error.
func (q *Queue[T]) Pop() (T, error) {
	if len(q.array) == 0 {
		return *new(T), errors.New("there is no elements in queue")
	}
	response := q.array[0]
	q.array = q.array[1:]
	return response, nil
}

// This method checks whether the queue is empty or not.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.array) == 0
}
func (q Queue[T]) Size() int {
	return len(q.array)
}

// This function creates and returns a new queue based on the given slice.
func NewQueue[T any](arr []T) *Queue[T] {
	return &Queue[T]{arr}
}
