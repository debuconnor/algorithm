package main

import (
	"fmt"
	"strconv"
)

func main() {
	var selfNum [10001]int

	for i := 1; i <= 10000; i++ {
		n := i
		str := strconv.Itoa(n)

		for _, v := range str {
			num, _ := strconv.Atoi(string(v))
			n += num
		}

		if n > 10000 {
			break
		}

		selfNum[n]++
	}

	for i, v := range selfNum {
		if v == 0 && i > 0 && i < 9994 {
			fmt.Println(i)
		}
	}
}
