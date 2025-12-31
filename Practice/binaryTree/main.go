package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)
// Walk 遍历树 t，并树中所有的值发送到信道 ch。
func Walk(t *tree.Tree, ch chan int) {
	var inorder func(*tree.Tree)
	inorder = func(n *tree.Tree) {
		if n == nil {
			return
		}
		inorder(n.Left)
		ch <- n.Value
		inorder(n.Right)
	}

	inorder(t)
	close(ch)
}

// Same 判断 t1 和 t2 是否包含相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	notSame := false
	for value1 := range ch1 {
		value2, ok := <- ch2
		if !((value1 == value2) && (ok)) {
			return false
		}
	}

	_, notSame = <- ch2
	return  !notSame
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
