package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := "Hello Woncome to JarJar"
	fmt.Println(a)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Got the input ", input)
}
