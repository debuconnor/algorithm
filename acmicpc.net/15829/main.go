package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	const FIX = 96
	var l, r, m, result uint64
	var str string

	r = 31
	m = 1234567891
	result = 0

	fmt.Fscan(reader, &l)
	fmt.Fscan(reader, &str)

	for i, v := range str{
		result += ((uint64(v) - FIX) * powInt(r, uint64(i))) % m
	}

	fmt.Fprintln(w, result)
	fmt.Fprintln(w, powInt(r, 49))
}

func powInt(x, y uint64) uint64 {
    return uint64(math.Pow(float64(x), float64(y)))
}