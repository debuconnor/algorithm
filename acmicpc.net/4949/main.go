package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var result []string

	for {
		str, _ := r.ReadString('\n')
		str = strings.ReplaceAll(str, "\n", "")
		if str == "." {
			break
		}

		var stack []string

		for _, v := range str{
			top := len(stack) - 1
			char := string(v)

			if top == -1 && (v == ')' || v == ']'){
				stack = append(stack, " ")
				break
			}

			if v == '(' || v == '['{
				stack = append(stack, char)
			} else if v == ')' {
				if stack[top] == "("{
					stack = stack[:top]
				}
			} else if v == ']' {
				if stack[top] == "["{
					stack = stack[:top]
				}
			}
		}

		if len(stack) == 0{
			result = append(result, "yes")
		} else {
			result = append(result, "no")
		}
	}

	for _, v := range result{
		fmt.Fprintln(w, v)
	}
}