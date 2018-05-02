package main

import (
	"fmt"
	"github.com/tzaffi/go-playground/z/node"
	"github.com/tzaffi/go-playground/z/q"
)

func main() {
	switch cmd := "BFS_BY_LEVEL"; cmd {
	case "BFS_BY_LEVEL":
		fmt.Println("BFS by level")
		bfsByLevel()
	default:
		fmt.Print("did not select a viable choice to run\n")
	}
}

/*  7
   / \
  8   4
 /\   /
1  7 3
  / \
 5   9  */
func getTreeExample() *node.Node {
	five, nine := node.Node{Val: 5}, node.Node{Val: 9}
	one, seven, three := node.Node{Val: 1}, node.Node{Val: 7, Left: &five, Right: &nine}, node.Node{Val: 3}
	eight, four := node.Node{Val: 8, Left: &one, Right: &seven}, node.Node{Val: 4, Left: &three}
	return &node.Node{Val: 7, Left: &eight, Right: &four}
}

func printBFS(root *node.Node) {
	if root == nil {
		return
	}
	currDepth := 0
	root.Depth = currDepth

	q := queue.NodeQ{}
	q.Push(root)
	for !q.Empty() {
		n, _ := q.Pop()
		if n.Depth > currDepth {
			fmt.Print("\n")
			currDepth = n.Depth
		}
		fmt.Printf("%d ", n.Val)
		l, r := n.Left, n.Right
		if l != nil {
			l.Depth = currDepth + 1
			q.Push(l)
		}
		if r != nil {
			r.Depth = currDepth + 1
			q.Push(r)
		}
	}
}

func bfsByLevel() {
	root := getTreeExample()
	printBFS(root)
}
