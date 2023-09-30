package main

import (
	"fmt"
	"math/rand"
)

func main() {
	header := "SpaceLine\t  Days\t  Trip-Type\t  Price\n"
	fmt.Print(header)
	for i := 0; i < len(header); i++ {
		fmt.Print("==")
	}
	var spaceLineSlice = []string{"Virgin Galactic", "SpaceX \t \t", "Space Adventures"}
	fmt.Print("\n")
	for i := 0; i < 10; i++ {
		var time int = rand.Intn(50)
		var price int = 0
		if time < 20 {
			price = rand.Intn(50-36) + 36
		} else {
			price = rand.Intn(35 - 0)
		}
		fmt.Printf("%v \t %d \t Round-trip \t $ %d \n", spaceLineSlice[rand.Intn(len(spaceLineSlice))], time, price)
	}

}
