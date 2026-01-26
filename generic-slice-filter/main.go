/*
The "Generic Slice Filter" 🧬
Go 1.18 introduced Generics, allowing us to write functions that work with any type while maintaining strict type safety. This is a massive improvement over using interface{} (now any).

📋 The Requirements:
The Function: Write a generic function called Filter. 🔍

The Signature:

It should take two arguments:

A slice of any type T.

A "predicate" function that takes a T and returns a bool.

It should return a new slice containing only the elements that passed the predicate.

The Types:

Use a Type Constraint so that T can be "any" type. 🧩

The Test: In your main() function, use the same Filter function to:

Filter a slice of integers to keep only even numbers.

Filter a slice of strings to keep only those with a length greater than 3. 🧵
*/
package main

import "fmt"

func Filter[T any](a []T, predicate func(T) bool) []T {
	var result []T

	for _, v := range a {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenInts := Filter(nums, func(num int) bool {
		return num%2 == 0
	})

	fmt.Println(evenInts)

	strs := []string{"my", "name", "is", "deepak", "dalal"}

	shortStrings := Filter(strs, func(str string) bool {
		return len(str) > 3
	})
	
	fmt.Println(shortStrings)
}