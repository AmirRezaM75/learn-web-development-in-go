package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3)) // 6
	numbers := []int{1, 2, 3}
	fmt.Println(sum(numbers...)) // 6
}

func sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
