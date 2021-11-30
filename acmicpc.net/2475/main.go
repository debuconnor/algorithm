package main

import "fmt"

func main() {
	var n [5]int
	result := 0

	for i, _ := range n {
		fmt.Scanf("%d", &n[i])

		n[i] *= n[i]
		result += n[i]
	}

	fmt.Println(result % 10)
}
