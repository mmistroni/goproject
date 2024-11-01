package main

import (
	"fmt"
	"myproject/mypackage"
)

func main() {
	person := mypackage.MyClass{
		Name: "Alice",
		Age:  30,
	}

	fmt.Println(person.Name, person.Age)
}
