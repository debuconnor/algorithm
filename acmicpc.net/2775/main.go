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

	var tc, n, k int

	fmt.Fscan(r, &tc)

	for t := 0; t < tc; t++{
		fmt.Fscan(r, &k, &n)

		apt := make([][]int, k + 1)
		
		for i := 0; i < k + 1; i++{
			for j := 0; j < n; j++{
				if j == 0{
					apt[i] = append(apt[i], 1)
				} else if i == 0{
					apt[i] = append(apt[i], j + 1)
				} else {
					apt[i] = append(apt[i], apt[i - 1][j] + apt[i][j - 1])
				}
			}
		}
	
		fmt.Fprintln(w, apt[k][n - 1])
	}

}