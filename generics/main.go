package main

import "fmt"

// type CustomeData interface {
// 	constraints.Ordered
// }

// type CustomeMap[T comparable] map[T]string

// type User[T CustomeData] struct {
// 	ID   int
// 	Name string
// 	Data T
// }

// type UserId int

// func Add[T constraints.Ordered](a, b T) T {
// 	return a + b
// }

// func subtract[T ~int | float64](a, b T) T {
// 	return a - b
// }

// func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
// 	var newValues []T
// 	for _, v := range values {
// 		newValue := mapFunc(v)
// 		newValues = append(newValues, newValue)
// 	}
// 	return newValues
// }

// func main() {

// 	u := User[int]{
// 		ID:   0,
// 		Name: "",
// 		Data: 10,
// 	}
// 	fmt.Println(u)

// 	a := UserId(1)
// 	b := UserId(3)
// 	subtract(a, b)
// 	fmt.Println(Add(3, 4))
// 	fmt.Println(Add(2.5, 1.5))

// 	result := MapValues([]int{1, 2, 3}, func(i int) int {
// 		return i * 5
// 	})
// 	fmt.Printf("result: %+v\n", result)
// }

func SlicesIndex[S []E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}

	}
	return -1
}

func main()  {
	var s = []int{1,2,3}
	i := SlicesIndex(s, 3)
	fmt.Println(i)

}
