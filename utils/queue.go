package utils

import "errors"

type queue struct {
	data []string
}

func NewQueue() *queue {
	return &queue{data: []string{}}
}

func (q *queue) Enqueue(s string) {
	q.data = append(q.data, s)
}

func (q *queue) Dequeue() (string, error) {
	if q.Empty() {
		return "", errors.New("antrian sudah kosong")
	}

	pop := q.data[0]
	q.data = q.data[1:]
	return pop, nil
}

func (q *queue) Front() (string, error) {
	if q.Empty() {
		return "", errors.New("antrian kosong")
	}
	return q.data[0], nil
}

func (q *queue) Empty() bool {
	return len(q.data) == 0
}

func (q *queue) Size() int {
	return len(q.data)
}
