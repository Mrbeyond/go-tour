package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	current := t

	if t == nil {
		return
	}

	if current.Left != nil {
		Walk(current.Left, ch)
	}
	ch <- current.Value

	Walk(current.Right, ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()

	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()

	for {
		v, ok1 := <-ch1
		j, ok2 := <-ch2

		// Check if either channel is closed
		if !ok1 || !ok2 {
			break
		}
		fmt.Printf("t1: %v, t2: %v\n", v, j)

		if v != j {
			return false
		}

	}

	return true
}

func main() {

	condition1 := Same(tree.New(1), tree.New(1))
	condition2 := Same(tree.New(1), tree.New(2))

	fmt.Println("Same(tree.New(1), tree.New(1)): ", condition1)
	fmt.Println("Same(tree.New(1), tree.New(2)): ", condition2)
}
