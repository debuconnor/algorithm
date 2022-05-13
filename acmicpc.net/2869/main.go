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

	var a, b, v int
	fmt.Fscan(r, &a, &b, &v)

	result := (v - a) / (a - b)
	
	if v - a > ((v - a) / (a - b)) * (a - b){
		fmt.Fprintln(w, result + 2)
	} else {
		fmt.Fprintln(w, result + 1)
	}
}