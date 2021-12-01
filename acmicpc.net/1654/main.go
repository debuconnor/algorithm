package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, n int
	var lan []int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &k, &n)

	temp := 0
	min := 2147483647
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &temp)
		lan = append(lan, temp)
	}

	fmt.Fprintln(writer, lan)
}
