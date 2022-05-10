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

	var arr, result []int

	for i := 1; i <= n; i++{
		arr = append(arr, i)
	}

	pointer := 0
	length := len(arr)

	for length > 0 {
		pointer += k - 1
		pointer %= length 

		result = append(result, arr[pointer])
		arr = append(arr[:pointer], arr[pointer + 1:]...)

		length-- 
	}

	length = len(result)

	fmt.Fprint(w, "<")
	for i, v := range result{
		fmt.Fprint(w, v)
		if i != length - 1{
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintln(w, ">")
}