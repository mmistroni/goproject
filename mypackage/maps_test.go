package mypackage

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {

	testMap := simpleFunction()

	if testMap[90] != "dog" {
		t.Errorf("Incorrect datqa details: %v", testMap)
	}
}

func TestUpdateMap(t *testing.T) {

	testMap := simpleFunction()

	testMap[11] = "GuineaPig"

	if testMap[90] == "" {
		t.Errorf("key 90 is busted: %v", testMap)
	}
}

func TestDeleteFromMap(t *testing.T) {

	testMap := simpleFunction()

	testMap[11] = "GuineaPig"

	if testMap[11] == "" {
		t.Errorf("key 90 is busted: %v", testMap)
	}

	delete(testMap, 11)

	animal, exists := testMap[11]

	if exists {
		t.Errorf(" is busted: %v", animal)
	} else {
		fmt.Println("the key has been deleted")
	}

}
