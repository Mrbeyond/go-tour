package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Works for RemoveFirst
//type List[T comparable] struct {
//	next *List[T]
//	val  T
//}

// Add to the list, return the list head and the total sze
func (l *List[T]) Add(data T) (*List[T], int) {

	newList := &List[T]{val: data}
	if l == nil {
		l = newList
		return l, 1
	}

	current := l
	ite := 1 // iteration count
	// get the last List with no next
	for current.next != nil {
		current = current.next
		ite++
	}
	current.next = newList

	return newList, ite + 1
}

func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

}

// Works with comparable
// func (l *List[T]) RemoveFirst(data T) {
// 	if l == nil {
// 		return
// 	}

// 	// If the head of the list has the value to be removed
// 	if l.val == data {
// 		*l = *l.next
// 		return
// 	}

// 	// Iterate through the list to find the value to remove
// 	current := l
// 	for current.next != nil {
// 		if current.next.val == data {
// 			current.next = current.next.next
// 			return
// 		}
// 		current = current.next
// 	}
// }

func main() {
	list := &List[string]{nil, "omo"}
	_, size := list.Add("ade")
  
	fmt.Println(fmt.Sprintf("Value: %v \t Size: %d", list, size))
  
	list.Add("ade1")
	list.Add("ade2")
	list.Add("ade3")
	list.Add("ade4")
	//list.RemoveFirst("omo") // works with comparable type
	list.Print()
}
