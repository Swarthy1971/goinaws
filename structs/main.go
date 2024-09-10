package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(newName string) {
	fmt.Println("address pf copy", &p.Name) // remove * before Person
	p.Name = newName
}

func main() {
	myPerson := NewPerson("Marek", 26)
	fmt.Println("address of allocated memory", &myPerson.Name)
	myPerson.changeName("John")

	fmt.Println(myPerson.Name)
	fmt.Println(myPerson.Age)
	fmt.Printf("this is my person %+v\n", myPerson)

	a := 7
	b := &a
	*b = 9
	fmt.Println(a, b)

	mySlice := []int{
		1, 2, 3,
	}
	for index, _ := range mySlice {
		mySlice[index]++
	}
	fmt.Println(mySlice)
}
