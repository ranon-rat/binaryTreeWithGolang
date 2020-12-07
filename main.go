package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) printNodes() {

	if n.left != nil {
		fmt.Println("left ", n.left)
		n.left.printNodes()
	}
	if n.right != nil {
		fmt.Println("right", n.right)
		n.right.printNodes()
	}
}

func (n *node) insert(number [][]int) {
	if len(number) >= 1 {
		if number[0][1] >= 1 {
			n.left = &node{val: number[0][0]}
			n.left.insert(number[1:])
		}
		if number[0][2] >= 1 {
			n.right = &node{val: number[0][0]}
			n.right.insert(number[1:])
		}
	}
}

func main() {
	var t *node = &node{}
	var n [][]int = [][]int{
		{1, 3, 1},
		{2, 2, 0},
		{3, 0, 3},
	}
	t.insert(n)
	t.printNodes()
	fmt.Printf("%#t\n", t)

}
