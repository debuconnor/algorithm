package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	const OPEN = 40
	const CLOSE = 41
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	var str string
	fmt.Fscanln(r, &n)

	count := 0

	for i := 0; i < n; i++{
		fmt.Fscan(r, &str)
		chars := []rune(str)
		length := len(chars)
		
		if length % 2 == 1{
			fmt.Fprintln(w, "NO")
			continue
		}
		
		for j, v := range chars{
			if v == OPEN{
				count++
			} else if v == CLOSE{
				count--
			}

			if count < 0{
				fmt.Fprintln(w, "NO")
				count = 0
				break
			} else if j == len(str) - 1{
				if count == 0{
					fmt.Fprintln(w, "YES")
				} else {
					fmt.Fprintln(w, "NO")
				}
				count = 0
			}
		}
	}
}