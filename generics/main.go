package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomeData interface {
	constraints.Ordered
}

type User[T CustomeData] struct {
	ID   int
	Name string
	Data T
}

type UserId int

func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

func subtract[T ~int | float64](a, b T) T {
	return a - b
}

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {

	u := User[int]{
		ID:   0,
		Name: "",
		Data: 10,
	}
	fmt.Println(u)

	a := UserId(1)
	b := UserId(3)
	subtract(a, b)
	fmt.Println(Add(3, 4))
	fmt.Println(Add(2.5, 1.5))

	result := MapValues([]int{1, 2, 3}, func(i int) int {
		return i * 5
	})
	fmt.Printf("result: %+v\n", result)
}
