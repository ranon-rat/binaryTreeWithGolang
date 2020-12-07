package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	val   int
	left  *node
	right *node
	d     int
}

func (n *node) printNodes() {
	n.d++
	if n.left != nil {
		fmt.Print("\033[31m")
		for i := 1; i < n.d; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println("left", n.left)
		n.left.d = n.d
		n.left.printNodes()
	}
	if n.right != nil {
		fmt.Print("\033[34m")
		for i := 1; i < n.d; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println("right", n.right)
		n.right.d = n.d
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
func randomNumbers(howMany int) [][]int {
	var num [][]int
	for len(num) < howMany {
		fmt.Print()
		num = append(num, []int{rand.Intn(15) - 10, rand.Intn(15+10) - 4, rand.Intn(15+10) - 4})
	}
	return num

}
func main() {
	var t *node = &node{}
	var n [][]int = randomNumbers(5)
	fmt.Println(n)
	t.insert(n)
	t.printNodes()
	fmt.Println("\033[0m")
	fmt.Printf("%#v\v", t)

}
