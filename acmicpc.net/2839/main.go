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
	fmt.Fscan(r, &n)

	result := -1

	for i := 0; i * 3 <= 12; i++{
		preCalc := n - i * 3
		
		if preCalc < 0 {
			break
		}
		
		if preCalc % 5 == 0{
			result = i
			result += preCalc / 5
		}
	}

	fmt.Fprintln(w, result)
}