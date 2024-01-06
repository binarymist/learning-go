// Taken from and expanded on https://go.dev/tour/generics/2

package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

// Add adds a new value of type T to the end of the list l.
func (list *List[T]) Add(val T) {
	for list.next != nil {
		list = list.next
	}
	list.next = &List[T]{val: val}
}

func (list *List[T]) Remove(val T) *List[T] {
	if list == nil {
		return nil
	}
	if list.val == val {
		return list.next
	}
	list.next = list.next.Remove(val)
	return list
}

func (list *List[T]) Contains(val T) bool {
	for list != nil {
		if list.val == val {
			return true
		}
		// Here we are not actually modifying the memory location that the original pointer points to.
		// Instead, we are modifying the local copy of the pointer to point to a different memory location.
		list = list.next
	}
	return false
}

func (list *List[T]) Size() int {
	count := 0
	for list != nil {
		count++
		list = list.next
	}
	return count
}

func (list *List[T]) PrintListVals() {
	current := list
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func textResponseIfContains(contains bool) (response string) {
	if contains {
		response = "contains"
	} else {
		response = "does not contain"
	}
	return
}

func main() {
	list := List[int]{}

	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	fmt.Println("The following is the contents of the list after attempting to Add: 1, 3, 5, 7")
	list.PrintListVals()

	list = *list.Remove(2)
	fmt.Println("The following is the contents of the list after attempting to Remove: 2")
	list.PrintListVals()

	list = *list.Remove(3)
	fmt.Println("The following is the contents of the list after attempting to Remove: 3")
	list.PrintListVals()

	fmt.Println(fmt.Sprintf("The list %s 1.", textResponseIfContains(list.Contains(1))))
	fmt.Println(fmt.Sprintf("The list %s 2.", textResponseIfContains(list.Contains(2))))
	fmt.Println(fmt.Sprintf("The list %s 3.", textResponseIfContains(list.Contains(3))))

	fmt.Println(fmt.Sprintf("The list size is %v now.", list.Size()))
}
