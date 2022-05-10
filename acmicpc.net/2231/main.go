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
	result := 0

	for i := 0; i < n; i++{
		if calc(i) == n {
			result = i
			break
		}
	}

	fmt.Fprintln(w, result)
}

func calc(num int) (result int){
	result = num

	for num >= 10{
		result += num % 10
		num /= 10
	}

	result += num

	return
}