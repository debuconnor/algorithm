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

	var x, y int
	
	fmt.Fscanf(r, "%d %d", &x, &y)
	
	for i := x; i <= y; i++{
		if i < 2{
			continue
		}

		if isPrimeNumber(i) {
			fmt.Fprintln(w, i)
		}
	}
}

func isPrimeNumber(num int) bool{
	for i := 2; i * i <= num; i++{
		if num % i == 0{
			return false
		}
	}

	return true
}