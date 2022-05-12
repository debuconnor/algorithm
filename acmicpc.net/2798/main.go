package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int

	fmt.Fscan(r, &n, &m)
	cards := make([]int, n)
	for i := 0; i < n; i++{
		fmt.Fscan(r, &cards[i])
	}

	sort.Ints(cards)
	sum := 0
	close := 999999

	for i := 0; i < n - 2; i++{
		for j := i + 1; j < n - 1; j++{

			for k := j + 1; k < n; k++{
				cardsum := cards[i] + cards[j] + cards[k]
				if cardsum > m{
					break
				}

				interval := abs(m - cardsum)

				if close > interval{
					close = interval
					sum = cardsum
				}
			}
		}
	}

	fmt.Fprintln(w, sum)
}

func abs(num int) int{
	if num < 0{
		return num - num * 2
	}

	return num
}