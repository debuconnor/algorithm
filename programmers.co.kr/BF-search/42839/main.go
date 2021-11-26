package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := "011"
	number, _ := strconv.Atoi(numbers)
	var origin = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	init := make([]int, 10)
	copy(init, origin)

	new := make([]int, 10)
	copy(new, origin)

	result := 0

	for i := 0; i < len(numbers); i++ {
		origin[numbers[i]-48]++
	}

	for i := 1; i <= number; i++ {
		temp := strconv.Itoa(i)

		for j := 0; j < len(temp); j++ {
			new[temp[j]-48]++
		}

		if isContain(origin, new) {
			if isPrimeNumber(i) {
				result++
			}
		}

		copy(new, init)
	}

	fmt.Println(result)
}

func isPrimeNumber(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func isContain(origin, new []int) bool {
	for i, v := range origin {
		if v-new[i] < 0 {
			return false
		}
	}

	return true
}
