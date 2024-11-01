package mypackage

import "testing"

func TestMyClass(t *testing.T) {
	person := MyClass{
		Name: "Bob",
		Age:  25,
	}

	if person.Name != "Bob" || person.Age != 25 {
		t.Errorf("Incorrect person details: %v", person)
	}
}
