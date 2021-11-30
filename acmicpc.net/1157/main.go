package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a [27]int
	s := ""

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscanln(reader, &s)

	len := len(s)
	for i := 0; i < len; i++ {
		char := s[i]
		if char >= 'a' {
			a[char-97]++
		} else {
			a[char-65]++
		}
	}

	temp := 0
	result := 0
	for i, v := range a {
		if temp < v {
			temp = v
			result = i
		}
	}

	count := 0
	for _, v := range a {
		if temp == v {
			count++
		}
	}

	if count > 1 {
		fmt.Fprintln(writer, "?")
	} else {
		fmt.Fprintln(writer, string(result+65))
	}
}
