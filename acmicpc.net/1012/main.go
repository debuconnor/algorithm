package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tc int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &tc)

	for i := 0; i < tc; i++ {

	}

	fmt.Fprintln(writer)
}
