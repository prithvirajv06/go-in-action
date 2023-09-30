package main

import (
	"fmt"
	"time"
)

func main() {

	timeNow := time.Now()
	fmt.Println("Jarjar its time ", timeNow)

	fmt.Println("Im not getting most out of it ")

	formattedTIme := timeNow.Format("01-02-2006 Monday")

	fmt.Println(formattedTIme)

	fmt.Println("Set the time line to 2000")

	t := time.Date(2000, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
