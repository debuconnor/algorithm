package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var fibonacci = [42]int{1, 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155}

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	temp := 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &temp)
		fmt.Fprintln(writer, fibonacci[temp], fibonacci[temp+1])
	}
}
