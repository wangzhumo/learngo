package impls

import "fmt"

func ExampleQueue_IsEmpty() {
	q := Queue{1, 2, 3}
	if q.IsEmpty() {
		fmt.Println("Empty")
	} else {
		fmt.Println("Not Empty")
	}

	// Output:
	// Not Empty
}
