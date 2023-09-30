package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Weolcome Jarja")
	fmt.Println("State your condition")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	convertion, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Print(error)
	} else {
		fmt.Println("Ok ill record this condition as ", convertion)
	}

}
