package main

import "fmt"

func main() {
	var arr [9]int
	maxIndex := -1
	maxNum := -1

	for i := 0; i < 9; i++ {
		fmt.Scanf("%d\n", &arr[i])
	}

	for i, v := range arr {
		if v > maxNum {
			maxNum = v
			maxIndex = i + 1
		}
	}

	fmt.Println(maxNum)
	fmt.Println(maxIndex)
}
