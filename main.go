package main

import (
	"fmt"
	"goproject/anotherpackage"
	"goproject/mypackage"
)


func main() {
	person := mypackage.MyClass{
		Name: "Alice",
		Age:  45,
	}

	another := anotherpackage.AnotherClass{
		Title:   "Mr",
		Sex:     "Male",
		Teacher: false,
	}

	fmt.Println(person.Name, person.Age)

	fmt.Println("--------------------")

	fmt.Println(another.Title, another.Teacher, another.Sex)

}
