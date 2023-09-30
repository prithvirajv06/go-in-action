package main

import "fmt"

//Public
const LoginToken string = "Some-Key"

func main() {

	var name string = "Prithviraj"
	fmt.Println(name)
	fmt.Printf("Type :%T", name)
	var numeric8 uint8 = 255
	fmt.Print(numeric8)
	fmt.Printf("Type :%T", numeric8)
	// Float 64 gives more pression then float 8

	// Implicit type
	var someName = "SomeName"
	fmt.Print(someName)

	// No  Var (Only inside method)
	someVar := "Some Var"
	fmt.Print(someVar)

	fmt.Println(LoginToken)
}
