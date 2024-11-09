// Sometimes we'll want to sort a collection by something
// other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts
// in Go.

package main

import (
	"fmt"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// We implement a comparison function for string
	// lengths. `cmp.Compare` is helpful for this.
	

	// Now we can call `slices.SortFunc` with this custom
	// comparison function to sort `fruits` by name length.
	
	fmt.Println(fruits)

	// We can use the same technique to sort a slice of
	// values that aren't built-in types.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// Sort `people` by age using `slices.SortFunc`.
	//
	// Note: if the `Person` struct is large,
	// you may want the slice to contain `*Person` instead
	// and adjust the sorting function accordingly. If in
	// doubt, [benchmark](testing-and-benchmarking)!
	
	fmt.Println(people)
}
