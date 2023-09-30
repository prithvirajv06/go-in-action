package main

import "fmt"

func main() {

	fmt.Println("Mapps")
	var mapsVar = make(map[string]string)
	mapsVar["name"] = "My Name"
	mapsVar["age"] = "20"

	delete(mapsVar, "age")

	fmt.Println(mapsVar)

	for key, value := range mapsVar {
		fmt.Println(key, value)
	}
}
