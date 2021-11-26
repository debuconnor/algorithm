package main

import (
	"fmt"
)

func main() {
	var progresses = []int{93, 30, 55}
	var speeds = []int{1, 30, 5}
	var day []int
	var result []int

	for i := 0; i < len(progresses); i++ {
		day = append(day, (100-progresses[i])/speeds[i])
		if (100-float32(progresses[i]))/float32(speeds[i]) > float32((100-progresses[i])/speeds[i]) {
			day[i]++
		}
	}

	temp := 0

	for i, v := range day {
		if i == 0 {
			if v < 1 {
				break
			}

			result = append(result, 1)
			temp = v
			continue
		}

		if v <= temp {
			result[len(result)-1]++
		} else {
			result = append(result, 1)
			temp = v
		}
	}

	fmt.Println(result)
}
