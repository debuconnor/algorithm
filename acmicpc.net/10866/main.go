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
	var deque []int
	fmt.Fscanln(r, &n)

	for i := 0; i < n; i++{
		fmt.Fscan(r, &cmd)
		if cmd == "push_front" || cmd == "push_back"{
			fmt.Fscan(r, &val)
		}

		length := len(deque)
		
		switch cmd{
		case "push_front":
			deque = append([]int{val}, deque...)

		case "push_back":
			deque = append(deque, val)

		case "pop_front":
			if length > 0{
				fmt.Fprintln(w, deque[0])
				deque = deque[1:]
			} else{
				fmt.Fprintln(w, -1)
			}

		case "pop_back":
			if length > 0{
				fmt.Fprintln(w, deque[length - 1])
				deque = deque[:length - 1]
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
				fmt.Fprintln(w, deque[0])
			} else{
				fmt.Fprintln(w, -1)
			}

		case "back":
			if length > 0{
				fmt.Fprintln(w, deque[length - 1])
			} else{
				fmt.Fprintln(w, -1)
			}
		}
	}
}