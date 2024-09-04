package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 4
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func activateGiftCard() func(int) int {
	amount := 100
	return func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
}

func main() {
	// res := plus(1, 2)
	// fmt.Println(res)

	// res = plusPlus(1, 2, 3)
	// fmt.Println(res)

	// a, b := vals()
	// fmt.Println(a)
	// fmt.Println(b)

	// _, c := vals()
	// fmt.Println(c)

	// sum(1, 2)
	// sum(1, 2, 3)

	// nums := []int{1, 2, 3, 4}
	// sum(nums...)

	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard1(15))
	fmt.Println(useGiftCard2(10))
}
