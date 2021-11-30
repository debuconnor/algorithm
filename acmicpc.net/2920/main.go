package main

import "fmt"

func main() {
	var note [8]int
	temp := 0
	result := ""

	for i, _ := range note {
		fmt.Scanf("%d", &note[i])

		if i == 0 {
			if note[i] == 1 {
				result = "ascending"
			} else if note[i] == 8 {
				result = "descending"
			} else {
				result = "mixed"
				break
			}
			temp = note[i]
		} else {
			if result == "ascending" {
				temp++
			} else if result == "descending" {
				temp--
			}

			if note[i] != temp {
				result = "mixed"
				break
			}
		}
	}

	fmt.Println(result)
}
