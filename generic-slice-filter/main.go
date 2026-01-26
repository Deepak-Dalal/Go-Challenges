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