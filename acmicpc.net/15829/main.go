package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var l, r, m, result uint64
	var str string

	r = 1
	m = 1234567891
	result = 0

	fmt.Fscan(reader, &l)
	fmt.Fscan(reader, &str)

	for _, v := range str{
		result = (result + (uint64(v - 'a') + 1) * r) % m
		r = (r * 31) % m
	}

	fmt.Fprintln(w, result)
}