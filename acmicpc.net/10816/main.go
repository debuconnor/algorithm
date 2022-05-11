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

	const NUMFIX = 10000000
	var n, m, tmp int
	matrix := make([]int, 20000001) // 0 = -10,000,000 / 10,000,000 = 0 / 10,000,000 = 20,000,001

	fmt.Fscan(r, &n)
	
	for i := 0; i < n; i++{
		fmt.Fscan(r, &tmp)
		matrix[tmp + NUMFIX]++
	}

	fmt.Fscan(r, &m)
	for i := 0; i < m; i++{
		fmt.Fscan(r, &tmp)
		fmt.Fprintf(w, "%d ", matrix[tmp + NUMFIX])
	}
}