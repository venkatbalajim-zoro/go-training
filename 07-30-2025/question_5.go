package day_one

import "fmt"

func sum(nums ...int) int {
	result := 0
	for _, val := range nums {
		result += val
	}
	return result
}

func VariadicFunc() {
	data := []int{1, 2, 3}
	fmt.Println(sum(data...))
}
