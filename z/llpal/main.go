package main

import "fmt"

func main() {
	/* leap years:
	fmt.Println(isLeapYear(2000) == true)
	fmt.Println(isLeapYear(2001) == false)
	fmt.Println(isLeapYear(2002) == false)
	fmt.Println(isLeapYear(2003) == false)
	fmt.Println(isLeapYear(2004) == true)
	fmt.Println(isLeapYear(2005) == false)
	fmt.Println(isLeapYear(2100) == false)
	*/
	s := "racecar"
	list := makeLL(s)
	fmt.Printf("%s : %t\n", s, isPal(list) == true)

	s = ""
	list = makeLL(s)
	fmt.Printf("%s : %t\n", s, isPal(list) == true)

	s = "donaldtrump"
	list = makeLL(s)
	fmt.Printf("%s : %t\n", s, isPal(list) == true)
}

func isLeapYear(y int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}

/* GLM: Palindrome. Write a function that checks if a linked list is a palindrome */

/* I. doubly linked list */
type Node struct {
	x    byte
	next *Node
}

func makeLL(s string) *Node {
	if len(s) == 0 {
		return nil
	}
	head := &Node{x: s[0], next: makeLL(s[1:])}
	return head
}

func getSlice(list *Node) []byte {
	a := []byte{}
	for c := list; c != nil; c = c.next {
		a = append(a, c.x)
	}
	return a
}

func isPal(list *Node) bool {
	a := getSlice(list)
	c := list
	for i := len(a) - 1; i >= 0; i-- {
		if c.x != a[i] {
			return false
		}
		c = c.next
	}
	return true
}
