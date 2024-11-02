package main

import (
	"fmt"
	"myproject/mypackage"
)

func main() {
	person := mypackage.MyClass{
		Name: "Alice",
		Age:  45,
	}

	fmt.Println(person.Name, person.Age)
}
