package main

import (
	"fmt"
)

func main() {
	var answers = []int{1, 3, 2, 4, 2}
	var stupid1 = [5]int{1, 2, 3, 4, 5}
	var stupid2 = [8]int{2, 1, 2, 3, 2, 4, 2, 5}
	var stupid3 = [10]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
	var score = [3]int{0, 0, 0}
	var result []int

	for i, v := range answers {
		if stupid1[i%len(stupid1)] == v {
			score[0]++
		}
		if stupid2[i%len(stupid2)] == v {
			score[1]++
		}
		if stupid3[i%len(stupid3)] == v {
			score[2]++
		}
	}

	highScore := highestNum(score)

	for i, v := range score {
		if v == highScore {
			result = append(result, i+1)
		}
	}

	fmt.Println(result)
}

func highestNum(arr [3]int) int {
	temp := 0
	for _, v := range arr {
		if temp < v {
			temp = v
		}
	}

	return temp
}
