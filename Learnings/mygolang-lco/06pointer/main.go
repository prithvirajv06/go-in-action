package main

import "fmt"

func main() {

	myNumber := 10
	myNumAddress := &myNumber

	fmt.Println("Address of My Number is  ", myNumAddress)
	fmt.Println("Value on the address is ", *myNumAddress)

	*myNumAddress = *myNumAddress + myNumber
	// Address value = Address value + variable value
	fmt.Println("Value on the address after aggr ", *myNumAddress)

}
