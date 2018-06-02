package main

import (
	"fmt"
	"github.com/tzaffi/go-playground/z/node"
	"github.com/tzaffi/go-playground/z/q"
)

func main() {
	switch cmd := "PAINT_FILL"; cmd {
	case "BFS_BY_LEVEL":
		fmt.Println("BFS by level")
		bfsByLevel()
	case "PAINT_FILL":
		fmt.Println("Paint Fill")
		paintFill()
	default:
		fmt.Print("did not select a viable choice to run\n")
	}
}

type Screen [][]uint32

var screen Screen
var w, h int

func (a Screen) String() string {
	res := ""
	for _, r := range a {
		res += fmt.Sprintf("%v\n", r)
	}
	return res
}

func paintFill() {
	w, h := 20, 10
	screen = Screen(make([][]uint32, h))
	for i := range screen {
		screen[i] = make([]uint32, w)
	}
	//make diagonal:
	for x, y := w-1, 0; x >= 0 && y < h; x, y = x-1, y+1 {
		screen[y][x] = 1
	}

	fmt.Printf("BEFORE:\n%s\n", screen.String())

	var paintR func(x, y int, oc, color uint32)
	paintR = func(x, y int, oc, color uint32) {
		if x < 0 || y < 0 || x >= w || y >= h || screen[y][x] != oc {
			return
		}
		screen[y][x] = color
		paintR(x+1, y, oc, color)
		paintR(x, y+1, oc, color)
		paintR(x-1, y, oc, color)
		paintR(x, y-1, oc, color)
	}

	paintIn := func(x, y int, color uint32) {
		oc := screen[y][x]
		paintR(x, y, oc, color)
	}

	paintIn(7, 7, 7)
	fmt.Printf("AFTER paintIn(%d, %d, %d):\n%s\n", 7, 7, 7, screen.String())

	paintIn(19, 9, 3)
	fmt.Printf("AFTER paintIn(%d, %d, %d):\n%s\n", 9, 19, 3, screen.String())
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
