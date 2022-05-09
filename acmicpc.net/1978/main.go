package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, tmp int
	num := make([]int, 1001)
	max := 0

	fmt.Fscan(r, &n)

	for i := 0; i < n; i++{
		fmt.Fscan(r, &tmp)
		if tmp > 1{
			num[tmp] = 1
		}

		if tmp > max{
			max = tmp
		}
	}

	num = getPrimeNumArray(max, num)

	count := 0

	for _, v := range num{
		if v == 1{
			count++
		}
	}

	fmt.Fprintln(w, count)
}

func getPrimeNumArray(numMax int, arr []int) []int{
	for i := 2; i <= numMax; i++{
		for j := 2; j * j <= i; j++{
			if i % j == 0{
				arr[i] = 0
			}
		}
	}
	return arr
}