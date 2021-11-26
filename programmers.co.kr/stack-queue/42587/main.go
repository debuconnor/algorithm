package main

import (
	"fmt"
)

func main() {
	var priorities = []int{2, 1, 3, 2}
	location := 2
	var q []int
	result := 0
	max := 0

	for _, v := range priorities {
		q = append(q, v)
		if v > max {
			max = v
		}
	}

	for {
		if q[0] == max {
			if location != 0 {
				q = q[1:]
				result++
				max = 0
				for _, v := range q {
					if v > max {
						max = v
					}
				}
			} else {
				break
			}
		} else if q[0] < max || q[0] < q[1] {
			q = append(q, q[0])
			q = q[1:]
		}

		if location == 0 {
			location = len(q) - 1
		} else {
			location--
		}
	}

	fmt.Println(result + 1)
}
