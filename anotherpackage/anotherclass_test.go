package anotherpackage

import "testing"

func TestMyClass(t *testing.T) {
	person := AnotherClass{
		Title:   "Mr",
		Sex:     "Male",
		Teacher: true,
	}

	if person.Teacher == false {
		t.Errorf("Incorrect person details: %v", person)
	}
}
