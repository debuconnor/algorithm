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

	var n, val int
	var cmd string
	var stack []int
	fmt.Fscanln(r, &n)

	for i := 0; i < n; i++{
		fmt.Fscan(r, &cmd)
		if cmd == "push"{
			fmt.Fscan(r, &val)
		}
		length := len(stack)

		switch cmd{
		case "push":
			stack = append(stack, val)

		case "pop":
			if length > 0{
				fmt.Fprintln(w, stack[length - 1])
				stack = stack[:length - 1]
			} else{
				fmt.Fprintln(w, -1)
			}

		case "size":
			fmt.Fprintln(w, length)

		case "empty":
			if length > 0{
				fmt.Fprintln(w, 0)
			} else {
				fmt.Fprintln(w, 1)
			}

		case "top":
			if length > 0{
				fmt.Fprintln(w, stack[length - 1])
			} else{
				fmt.Fprintln(w, -1)
			}
		}
	}
}