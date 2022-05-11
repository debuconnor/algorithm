package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var tc, h, w, n int

	fmt.Fscan(r, &tc)

	for i := 0; i < tc; i++{
		fmt.Fscan(r, &h)
		fmt.Fscan(r, &w)
		fmt.Fscan(r, &n)

		x := n / h + 1
		y := n % h

		if n % h == 0{
			y = h
			x -= 1
		}

		if n <= h {
			x = 1
		}

		fmt.Fprintf(wr, "%d%02d\n", y, x)
	}
}