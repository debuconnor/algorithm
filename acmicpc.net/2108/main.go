package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, tmp int
	var freq [8001]int // -4000 = 0, 0 = 4000, 4000 = 8001
	var mid []int
	avg := 0.0
	min := 4001
	max := -4001

	fmt.Fscan(r, &n)
	
	for i := 0; i < n; i++{
		fmt.Fscan(r, &tmp)
		avg += float64(tmp)
		freq[tmp + 4000]++
		mid = append(mid, tmp)

		if tmp < min{
			min = tmp
		}

		if tmp > max{
			max = tmp
		}
	}

	sort.Ints(mid)
	
	fmt.Fprintln(w, int(math.Round(avg / float64(n))))
	fmt.Fprintln(w, mid[len(mid) / 2])
	fmt.Fprintln(w, getMode(freq))
	fmt.Fprintln(w, max - min)
}

func getMode(arr [8001]int) int{
	max := 0
	count := 0
	result := 9999

	for i, v := range arr{
		if max <= v{
			max = v
			count++
			result = i - 4000
		}
	}

	if count > 1{
		flag := false
		for i, v := range arr{
			if v == max{
				if flag{
					return i - 4000
				}

				flag = true
			}
		}
	}

	return result
}