package main

import (
	"fmt"
)

func main() {
	var nums [10]int
	temp := make(map[int]bool)
	result := 0

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d\n", &nums[i])
	}

	for _, v := range nums {
		if temp[v%42] {
			continue
		}

		temp[v%42] = true
		result++
	}

	fmt.Println(result)
}
