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

	var n int
	fmt.Fscanln(r, &n)

	count := 0
	max := 1
	min := 0

	for {
		max += count * 6

		if min < n && n <= max{
			fmt.Fprintln(w, count + 1)
			break
		}
		count++
		min = max
	}
}