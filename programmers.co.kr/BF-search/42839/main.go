package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := "177"
	var origin = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	init := make([]int, 10)
	copy(init, origin)

	new := make([]int, 10)
	copy(new, origin)

	result := 0

	for i := 0; i < len(numbers); i++ {
		origin[numbers[i]-48]++
	}

	for i := 1; i <= maxNumber(origin); i++ {
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
	if num == 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func maxNumber(numArr []int) int {
	num := ""

	for i := 9; i >= 0; i-- {
		_num := strconv.Itoa(i)

		for j := numArr[i]; j > 0; j-- {
			num += _num
		}
	}

	maxNumber, _ := strconv.Atoi(num)
	return maxNumber
}

func isContain(origin, new []int) bool {
	for i, v := range origin {
		if v-new[i] < 0 {
			return false
		}
	}

	return true
}
