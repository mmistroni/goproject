package mypackage

func sum(numbers ...int) int {
	thesum := 0
	for _, item := range numbers {
		thesum += item
	}
	return thesum
}
