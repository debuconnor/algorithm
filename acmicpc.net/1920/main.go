package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, tmp int
	var a, b []int

	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++{
		fmt.Fscan(reader, &tmp)
		a = append(a, tmp)	
	}

	fmt.Fscan(reader, &m)
	for i := 0; i < m; i++{
		fmt.Fscan(reader, &tmp)
		b = append(b, tmp)	
	}

	sort.Ints(a)

	result := make([]int, m)

	for i, v := range b{
		left := 0
		right := n - 1

		for {
			mid := (left + right) / 2
			
			if v == a[mid]{
				result[i]++
				break
			} else if v > a[mid]{
				left = mid + 1
			} else if v < a[mid]{
				right = mid - 1
			}

			if left > right{
				break
			}
		}
	}

	for _, v := range result{
		fmt.Fprintln(writer, v)
	}
}