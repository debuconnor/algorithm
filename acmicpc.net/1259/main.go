package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := ""

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		fmt.Fscanln(reader, &s)

		if s == "0" {
			break
		}

		len := len(s)
		char := []byte(s)
		result := "no"

		switch len {
		case 1:
			result = "yes"
		case 2:
			if char[0] == char[1] {
				result = "yes"
			}
		case 3:
			if char[0] == char[2] {
				result = "yes"
			}
		case 4:
			if char[0] == char[3] {
				if char[1] == char[2] {
					result = "yes"
				}
			}
		case 5:
			if char[0] == char[4] {
				if char[1] == char[3] {
					result = "yes"
				}
			}
		}

		fmt.Fprintln(writer, result)
	}
}
