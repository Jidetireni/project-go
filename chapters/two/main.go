package main

import "fmt"

func linear_search(array []int, search_value int) int {
	for index, element := range array {
		if element == search_value {
			return index
		} else if element > search_value {
			break
		}
	}
	return -1
}

func binary_search(array []int, search_value int) int {
	lower_bound := 0
	upper_bound := len(array) - 1

	for lower_bound <= upper_bound {
		midpoint := (upper_bound + lower_bound) / 2
		value_at_midpoint := array[midpoint]

		if search_value == value_at_midpoint {
			return midpoint
		} else if search_value < value_at_midpoint {
			upper_bound = midpoint - 1
		} else {
			lower_bound = midpoint + 1
		}
	}
	return -1
}

func main() {
	array := []int{3, 17, 75, 80, 202}
	search_value := 22

	result := linear_search(array, search_value)
	result2 := binary_search(array, search_value)
	if result != -1 || result2 != -1 {
		if result != -1 {
			fmt.Printf("Value %d found at index %d using linear search\n", search_value, result)
		}
		if result2 != -1 {
			fmt.Printf("Value %d found at index %d using binary search\n", search_value, result)
		}
	} else {
		fmt.Printf("Value %d not found in the array\n", search_value)

	}
}
