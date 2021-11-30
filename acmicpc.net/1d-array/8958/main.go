package main

import (
	"fmt"
)

func main() {
	var tc int
	var temp byte

	fmt.Scanf("%d\n", &tc)
	result := make([]int, tc)
	s := ""

	for i := 0; i < tc; i++ {
		fmt.Scanln(&s)

		temp = 'Z'
		_result := 0
		for j := 0; j < len(s); j++ {
			if s[j] == 'O' {
				result[i]++
				if temp == s[j] {
					_result++
				}
			} else {
				_result = 0
			}

			result[i] += _result
			temp = s[j]
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
