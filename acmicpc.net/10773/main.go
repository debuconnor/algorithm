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

	var stack []int
	var k, tmp int
	result := 0

	fmt.Fscanln(r, &k)

	for ; k > 0; k--{
		fmt.Fscanln(r, &tmp)

		if tmp > 0 {
			result += tmp
			stack = append(stack, tmp)	
		} else {
			result -= stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}		
	}

	fmt.Fprintln(w, result)
}