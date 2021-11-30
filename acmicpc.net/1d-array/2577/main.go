package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums [10]int
	result := 1

	for i := 0; i < 3; i++ {
		temp := 0
		fmt.Scanf("%d\n", &temp)
		result *= temp
	}

	_result := strconv.Itoa(result)

	for i := 0; i < len(_result); i++ {
		nums[_result[i]-48]++
	}

	for _, v := range nums {
		fmt.Println(v)
	}
}
