package main

func bubble_sort(list []int) []int {
	unsorted_until_index := len(list)
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < unsorted_until_index; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sorted = false
			}
		}
		unsorted_until_index -= 1
	}
	return list
}
