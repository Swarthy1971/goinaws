package main

import "fmt"

func main() {
	myName := "Marek"
	myInt := 10
	myFloat := 10.5
	//var myName string = "Marek"
	fmt.Printf("Hello, my name is %s my int is %d my float is %f\n", myName, myInt, myFloat)

	var myFriendsName string
	var myBool bool
	var myOtherInt int
	myFriendsName = "John"
	myBool = true
	myOtherInt = 100

	fmt.Printf("my other frinds name %s my bool %t my other int %d\n", myFriendsName, myBool, myOtherInt)

}
