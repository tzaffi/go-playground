package main

import (
	"fmt"
	"github.com/tzaffi/go-playground/z/node"
	"github.com/tzaffi/go-playground/z/q"
)

func main() {
	switch cmd := "SORT_STACK"; cmd {
	case "BFS_BY_LEVEL":
		fmt.Println("BFS by level")
		bfsByLevel()
	case "PAINT_FILL":
		fmt.Println("Paint Fill")
		paintFill()
	case "ONE_AWAY":
		fmt.Println("One Away")
		oneAwayMain()
	case "BINODE":
		fmt.Println("Binode")
		binodeMain()
	case "SORT_STACK":
		fmt.Println("Sort Stack")
		sortStackMain()
	default:
		fmt.Print("did not select a viable choice to run\n")
	}
}

/********** SORT_STACK  ********/

func sortStackMain() {
	myStack := Stack([]int{1, 2, 3, 4, 5})
	fmt.Printf("myStack.Peek() = %d\n", myStack.Peek())

	fmt.Printf("myStack.Pop() = %d\n", myStack.Pop())
	fmt.Printf("myStack.Peek() = %d\n", myStack.Peek())

	myStack.Push(17)
	fmt.Printf("myStack.Peek() = %d\n", myStack.Peek())

	fmt.Printf("myStack = %s\n", &myStack)

	myStack = Stack([]int{5, 2, 1, 4, 7, 6})
	fmt.Printf("BEFORE: %s\n", &myStack)
	Sort(&myStack)
	fmt.Printf("AFTER: %s\n", &myStack)
	for ; !myStack.Empty(); myStack.Pop() {
		fmt.Println(myStack.Peek())
	}
}

func Sort(s *Stack) {
	if s == nil {
		return
	}
	sorted := Stack([]int{})
	sortR(s, &sorted)
	//	fmt.Printf("sort: s, sorted = %s, %s\n", s, &sorted)
	*s = sorted
}

func sortR(s, sorted *Stack) {
	//	fmt.Printf("sortR: s, sorted = %s, %s\n", s, sorted)
	if s.Empty() {
		return
	}
	v := s.Pop()
	var numPops int
	for numPops = 0; !sorted.Empty() && sorted.Peek() < v; numPops++ {
		s.Push(sorted.Pop())
	}
	sorted.Push(v)
	for i := 0; i < numPops; i++ {
		sorted.Push(s.Pop())
	}
	sortR(s, sorted)
}

type Stack []int

func (s *Stack) String() string {
	return fmt.Sprintf("%v", []int(*s))
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() int {
	if s.Empty() {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	if s.Empty() {
		return 0
	}
	pop := s.Peek()
	*s = append((*s)[0 : len(*s)-1])
	return pop
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

/********** BI_NODE  ********/

func binodeMain() {
	fmt.Printf("[0, 4, 5] --> %s\n", slice2dll([]int{0, 4, 5}))
	fullTree := makeFullTree()
	fullTreeBFS := binode2bfs(fullTree)
	printBFS(fullTreeBFS)

	testCases := []binodeTC{
		binodeTC{
			root:        fullTree,
			expectedDll: slice2dll([]int{1, 2, 3, 4, 5, 7, 8, 9, 11}),
		},
		binodeTC{
			root:        nil,
			expectedDll: nil,
		},
		binodeTC{
			root:        &Binode{data: 17},
			expectedDll: &Binode{data: 17},
		},
		binodeTC{
			root:        makeLeftTree(),
			expectedDll: slice2dll([]int{1, 2, 3, 5}),
		},
	}

	for _, tc := range testCases {
		actualDll, _ := listify(tc.root)
		actStr := actualDll.String()
		expStr := tc.expectedDll.String()
		fmt.Printf("expected: %s VS actual: %s\n", expStr, actStr)
		if actStr == expStr {
			fmt.Println("SUCCESS!!!!!!!!!!")
		} else {
			fmt.Println("FAIL :-(")
		}
	}
}

func listify(root *Binode) (min, max *Binode) {
	if root == nil {
		return
	}

	min, lhead := listify(root.n1)
	rtail, max := listify(root.n2)
	if lhead == nil {
		min = root
	} else {
		lhead.n2 = root
		root.n1 = lhead
	}
	if rtail == nil {
		max = root
	} else {
		rtail.n1 = root
		root.n2 = rtail
	}
	return
}

func (dll *Binode) String() string {
	if dll == nil {
		return ""
	}
	return fmt.Sprintf("[%d]-%s", dll.data, dll.n2.String())
}

type Binode struct {
	n1, n2 *Binode
	data   int
}

type binodeTC struct {
	root, expectedDll *Binode
}

/*      5
     /    \
    3      9
   / \    / \
  2   4  7  11
 /        \
1          8 */
/* --> [1, 2, 3, 4, 5, 7, 8, 9, 11] */
func makeFullTree() *Binode {
	one := &Binode{data: 1}
	two := &Binode{data: 2, n1: one}
	four := &Binode{data: 4}
	three := &Binode{data: 3, n1: two, n2: four}
	seven := &Binode{data: 7, n2: &Binode{data: 8}}
	nine := &Binode{data: 9, n1: seven, n2: &Binode{data: 11}}
	return &Binode{data: 5, n1: three, n2: nine}
}

/*      5
     /
    3
   /
  2
 /
1   */
/* --> [1, 2, 3, 5] */
func makeLeftTree() *Binode {
	one := &Binode{data: 1}
	two := &Binode{data: 2, n1: one}
	three := &Binode{data: 3, n1: two}
	return &Binode{data: 5, n1: three}
}

func binode2bfs(root *Binode) *node.Node {
	if root == nil {
		return nil
	}
	return &node.Node{
		Val:   root.data,
		Left:  binode2bfs(root.n1),
		Right: binode2bfs(root.n2),
	}
}

func slice2dll(s []int) *Binode {
	if len(s) == 0 {
		return nil
	}
	bn := &Binode{
		n2:   slice2dll(s[1:]),
		data: s[0],
	}
	if bn.n2 != nil {
		bn.n2.n1 = bn
	}
	return bn
}

/********** ONE_AWAY ********/

type oneAwayTC struct {
	x, y   string
	expect bool
}

func oneAwayMain() {
	tcs := []oneAwayTC{
		oneAwayTC{"pale", "ple", true},
		oneAwayTC{"pales", "pale", true},
		oneAwayTC{"pale", "bale", true},
		oneAwayTC{"pale", "palk", true},
		oneAwayTC{"pale", "bake", false},
		oneAwayTC{"pale", "ale", true},
		oneAwayTC{"pale", "pal", true},
		oneAwayTC{"p", "", true},
		oneAwayTC{"", "p", true},
		oneAwayTC{"", "", true},
		oneAwayTC{"", "abc", false},
		oneAwayTC{"pale", "pale", true},
	}
	for _, tc := range tcs {
		x, y := tc.x, tc.y
		fmt.Printf("%s vs %s == %t ? --> %t\n", x, y, tc.expect, oneAway(x, y))
	}
}

func oneAway(x, y string) bool {
	if len(x) == len(y) {
		return sameOrReplace1(x, y)
	}
	if len(x) < len(y) {
		x, y = y, x
	}
	if len(x) > len(y)+1 {
		return false
	}
	if len(x) == 1 {
		return true
	}
	var i int
	for i = 0; i < len(y) && x[i] == y[i]; i++ {
	}
	return x[i+1:] == y[i:]
}

func sameOrReplace1(x, y string) bool {
	var i int
	for i = 0; i < len(y) && x[i] == y[i]; i++ {
	}
	if i == len(x) {
		return true
	}
	return x[i+1:] == y[i+1:]
}

/********** PAINT_FILL *********/

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
