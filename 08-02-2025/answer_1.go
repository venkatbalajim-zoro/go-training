package answers

import "fmt"

func Answer1() {
	var target int
	var length int

	fmt.Print("Enter the digit: ")
	fmt.Scan(&target)
	fmt.Print("Enter the length: ")
	fmt.Scan(&length)

	fmt.Print("Enter the elements: ")
	var array []int
	for i := 0; i < length; i++ {
		var elt int
		fmt.Scan(&elt)
		array = append(array, elt)
	}

	var indices []int
	for idx, elt := range array {
		temp := elt
		for temp > 0 {
			digit := temp % 10
			if digit == target {
				indices = append(indices, idx)
				break
			} else {
				temp /= 10
			}
		}
	}

	answer := 0
	for _, elt := range indices {
		answer += elt
	}

	fmt.Println("Answer is", answer)
}
