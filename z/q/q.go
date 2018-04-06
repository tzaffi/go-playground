package queue

import (
	"fmt"
)

type Q struct {
	q []int
}

func (q *Q) Empty() bool {
	return len(q.q) == 0
}

func (q *Q) Push(x int) {
	q.q = append(q.q, x)
}

func (q *Q) Pop() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("cannot Pop() an empty queue")
	}
	res := q.q[0]
	q.q = q.q[1:]
	return res, nil
}
