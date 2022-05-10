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
	var queue []int
	fmt.Fscanln(r, &n)

	for i := 0; i < n; i++{
		fmt.Fscan(r, &cmd)
		if cmd == "push"{
			fmt.Fscan(r, &val)
		}

		length := len(queue)
		
		switch cmd{
		case "push":
			queue = append(queue, val)

		case "pop":
			if length > 0{
				fmt.Fprintln(w, queue[0])
				queue = queue[1:]
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

		case "front":
			if length > 0{
				fmt.Fprintln(w, queue[0])
			} else{
				fmt.Fprintln(w, -1)
			}

		case "back":
			if length > 0{
				fmt.Fprintln(w, queue[length - 1])
			} else{
				fmt.Fprintln(w, -1)
			}
		}
	}
}