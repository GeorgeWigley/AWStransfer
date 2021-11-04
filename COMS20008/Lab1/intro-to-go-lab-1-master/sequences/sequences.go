package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	o := []int{}
	for k := range slice {
		o = append(o, f(slice[k]))
	}
	fmt.Println(o)
}

func mapArray(f func(a int) int, array [3]int) {
	o := []int{}
	for k := range array {
		o = append(o, f(array[k]))
	}
	fmt.Println(o)
}

func main() {
	intss := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intss)
	intsa := [3]int{1, 2, 3}
	mapArray(addOne, intsa)

	newSlice := intss[1:3]
	mapSlice(square, newSlice)

	fmt.Println(double(newSlice))
}
