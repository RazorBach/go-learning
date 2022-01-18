package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// Note: use a closure to make this function neat and close the channel easily
func Walk(t *tree.Tree, ch chan int) {
	var worker func(t *tree.Tree)
	worker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		worker(t.Left)
		ch <- t.Value
		worker(t.Right)
	}
	worker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
	return true
}

func Print(t *tree.Tree) {
	ch := make(chan int)
	go Walk(t, ch)
	for {
		v, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	Print(tree.New(1))
	Print(tree.New(2))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
