package queue

import (
	"fmt"

	"github.com/tzaffi/go-playground/z/node"
)

type IntQ struct {
	q []int
}

func (q *IntQ) Empty() bool {
	return len(q.q) == 0
}

func (q *IntQ) Push(x int) {
	q.q = append(q.q, x)
}

func (q *IntQ) Pop() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("cannot Pop() an empty queue")
	}
	res := q.q[0]
	q.q = q.q[1:]
	return res, nil
}

// --------- node queue's
type NodeQ struct {
	q []*node.Node
}

func (q *NodeQ) Empty() bool {
	return len(q.q) == 0
}

func (q *NodeQ) Push(x *node.Node) {
	q.q = append(q.q, x)
}

func (q *NodeQ) Pop() (*node.Node, error) {
	if q.Empty() {
		return nil, fmt.Errorf("cannot Pop() an empty queue")
	}
	res := q.q[0]
	q.q = q.q[1:]
	return res, nil
}
