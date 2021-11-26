package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []int{67, 676, 677}
	result := ""
	ssort(&numbers, 0, 1)

	for _, v := range numbers {
		if result == "" && v == 0 {
			continue
		}

		result += strconv.Itoa(v)
	}

	if result == "" {
		result = "0"
	}

	fmt.Println(result)
}

func ssort(arr *[]int, start int, pointer int) {
	if start >= len(*arr) || start >= pointer {
		return
	}

	if pointer >= len(*arr) {
		ssort(arr, start+1, start+2)
		return
	}

	x := strconv.Itoa((*arr)[start])
	y := strconv.Itoa((*arr)[pointer])

	if compare(x, y) {
		(*arr)[start], (*arr)[pointer] = (*arr)[pointer], (*arr)[start]
	}
	ssort(arr, start, pointer+1)
}

func compare(num1 string, num2 string) bool {
	temp1, _ := strconv.Atoi(num1 + num2)
	temp2, _ := strconv.Atoi(num2 + num1)

	return temp1 < temp2
}
