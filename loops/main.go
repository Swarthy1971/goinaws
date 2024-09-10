package main

import (
	"fmt"
	"slices"
)

func main() {
	//animals := [2]string{
	//	"cat",
	//	"dog"
	//}

	//animals[0] = "dog"
	//animals[1] = "cat"
	//fmt.Println(animals)

	// slices
	animals := []string{
		"dog",
		"cat",
	}
	animals = append(animals, "bird")
	animals = slices.Delete(animals, 0, 1)
	fmt.Println(animals)

}
