package main

import (
	"fmt"
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
	//animals = slices.Delete(animals, 0, 1)
	fmt.Println(animals)

	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	for index, value := range animals {
		fmt.Printf("index: %d value: %s\n", index, value)
	}
	for value := range 10 {
		fmt.Println(value)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
