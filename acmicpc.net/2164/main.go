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

	var matrix, matrix2 []int

	for i := 1; i <= n; i++{
		matrix = append(matrix, i)
	}

	length := n

	for length > 1 {
		for i, v := range matrix{
			if i == length - 1 && i % 2 == 0{
				matrix2 = append([]int{v}, matrix2...)
			} else if i % 2 == 1 {
				matrix2 = append(matrix2, v)
			}
		}
		matrix = matrix[:len(matrix2)]
		copy(matrix, matrix2)
		matrix2 = matrix2[:0]
		length = len(matrix)
	}

	
	fmt.Fprintln(w, matrix[0])
}