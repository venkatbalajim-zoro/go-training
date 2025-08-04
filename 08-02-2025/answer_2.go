package answers

import (
	"fmt"
	"math"
)

func Answer2() {
	var length int

	fmt.Print("Enter the length: ")
	fmt.Scan(&length)

	fmt.Print("Enter the elements: ")
	var array []int
	for i := 0; i < length; i++ {
		var elt int
		fmt.Scan(&elt)
		array = append(array, elt)
	}

	answer := math.MinInt

	for i := 0; i < length-1; i++ {
		elt1, elt2 := array[i], array[i+1]
		sum := elt1 + elt2
		if sum > answer {
			answer = sum
		}
	}

	fmt.Println("Answer is", answer)
}
