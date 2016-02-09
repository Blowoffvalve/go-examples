package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values from the tree to channel ch.

func WalkInOrder(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		WalkInOrder(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkInOrder(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	// Cannot use nest function declaration.
	WalkInOrder(t, ch)
	// The channel should be closed.
	close(ch)
}

// Same determines whether the trees t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for e1 := range ch1 {
		e2, ok := <-ch2
		if !(ok && e1 == e2) {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(10), tree.New(10)))
}
