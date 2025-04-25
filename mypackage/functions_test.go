package mypackage

import (
	"testing"
)

func TestSum(t *testing.T) {

	expectedSum := 4 * 5 / 2

	res := sum(1, 2, 3, 4)

	if res != expectedSum {
		t.Errorf("Incorrect result. Expected %v but was %v", expectedSum, res)
	}
}
