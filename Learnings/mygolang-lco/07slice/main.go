package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice = []string{"Gova", "Peach"}

	fmt.Println(mySlice)
	mySlice = append(mySlice, "Apple", "Orange")

	fmt.Println("After Append ", mySlice)

	fmt.Println(mySlice[1:3])

	var memoriedSlice = make([]string, 2)

	fmt.Println(memoriedSlice)
	//memoriedSlice[3] Will give error memeory out of bound
	memoriedSlice = append(memoriedSlice, mySlice...)

	fmt.Println("M-Sclice: ", memoriedSlice)

	//sort.Ints(memoriedSlice) if int and float
	sort.StringSlice(memoriedSlice).Sort()

	fmt.Println(memoriedSlice)

	var numbers = []int{0, 1, 2, 3, 4, 5, 6}
	numbers = append(numbers[:3], numbers[4:]...)
	fmt.Println(numbers)
}
