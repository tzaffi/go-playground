package queue

import (
	"fmt"
	"testing"
)

func TestIntQueue(t *testing.T) {
	q := IntQ{}

	if !q.Empty() {
		t.Error("q should be empty")
	}

	q.Push(3)
	if q.Empty() {
		t.Error("q should not be empty")
	}

	if x, e := q.Pop(); e != nil {
		t.Error("no error should result when popping non-empty")
	} else {
		if x != 3 {
			t.Error("should have retrieved the thing that we just pushed")
		}
		if !q.Empty() {
			t.Error("q should be empty again")
		}
	}

	q = IntQ{}
	ordinals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, x := range ordinals {
		fmt.Printf("ordinal = %d\n", x)
		q.Push(x)
		if q.Empty() {
			t.Errorf("q should not be empty after %d insertions", x)
		}
	}

	popped := []int{}
	var p int
	var e error
	for !q.Empty() {
		p, e = q.Pop()
		fmt.Printf("popped = %d\n", p)
		if e != nil {
			t.Errorf("should have gotten no error but got %s", e)
			continue
		}
		popped = append(popped, p)
	}
	fmt.Printf("popped = %v\n", popped)

	if len(ordinals) != len(popped) {
		t.Errorf("len(ordinals) != len(popped) because %d != %d", len(ordinals), len(popped))
	}

	for i := range ordinals {
		if ordinals[i] != popped[i] {
			t.Errorf("ordinals[i] != popped[i] because %d != %d", ordinals[i], popped[i])
		}
	}

}
