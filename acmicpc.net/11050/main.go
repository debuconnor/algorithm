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

	var n, k int

	fmt.Fscanln(r, &n, &k)

	result := getFactorial(n) / (getFactorial(k) * getFactorial(n - k))

	fmt.Fprintln(w, result)
}

func getFactorial(num int) (result int){
	result = 1

	for i := 1; i <= num; i++{
		result *= i
	}

	return
}