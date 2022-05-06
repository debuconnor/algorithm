package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := 0
	fmt.Fscanln(reader, &n)

	var input []int

	for i := 0; i < n; i++{
		tmp := 0
		fmt.Fscanln(reader, &tmp)
		input = append(input, tmp)
	}

	var stack []int
	var result []string
	top := 0
	count := 1
	moved := 0
	inputPointer := 0

	for {
		if top < 0 || top >= n || inputPointer >= n{
			break
		}

		stack = push(stack, &top, count)
		result = append(result, "+")
		// fmt.Println("\n>>push count: ", count, " moved:", moved, " pointer: ", inputPointer, " top: ", top)

		if stack[top - 1] < input[inputPointer]{
			count++
			continue
		}

		for {
			if top == 0 || inputPointer > n{
				count += moved + 1
				moved = 0
				break
			}

			if stack[top - 1] == input[inputPointer]{
				stack = pop(stack, &top)
				result = append(result, "-")
				count--
				moved++
				inputPointer++

				// fmt.Println("<<pop  count: ", count, " moved:", moved, " pointer: ", inputPointer, " top: ", top)
			} else {
				count += moved + 1
				moved = 0
				break				
			}
		}
	}

	if len(stack) == 0{
		for _, v := range result{
			fmt.Fprintln(writer, v)
		}
	} else {
		fmt.Println("NO")
	}
}

func push(arr []int, top *int, val int) []int{
	arr = append(arr, val)
	*top++

	return arr
}

func pop(arr []int, top *int) []int{
	arr = arr[:*top - 1]
	*top--

	return arr
}