package main

import (
	"fmt"
)

func main() {
	var tc int
	var c int

	fmt.Scanf("%d\n", &tc)

	result := make([]float64, tc)

	for i := 0; i < tc; i++ {
		fmt.Scanf("%d", &c)

		student := make([]float64, c)
		avg := 0.0

		for j := 0; j < c; j++ {
			fmt.Scanf("%f", &student[j])
			avg += student[j]
		}

		avg /= float64(c)
		success := 0

		for _, w := range student {
			if avg < w {
				success++
			}
		}

		result[i] = float64(success) / float64(c)
	}

	for _, v := range result {
		fmt.Printf("%.3f%%\n", v*100)
	}
}
