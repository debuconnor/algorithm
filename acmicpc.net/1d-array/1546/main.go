package main

import (
	"fmt"
)

func main() {
	var tc int

	fmt.Scanf("%d\n", &tc)
	subject := make([]float64, tc)
	maxScore := 0.0
	avg := 0.0

	for i := 0; i < tc; i++ {
		fmt.Scanf("%f", &subject[i])
	}

	for _, v := range subject {
		if maxScore < v {
			maxScore = v
		}
	}

	for i, v := range subject {
		subject[i] = v / maxScore * 100
		avg += subject[i]
	}

	fmt.Println(avg / float64(tc))
}
