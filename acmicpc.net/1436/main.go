package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sss []int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; ; i++ {
		if len(sss) == 10001 {
			break
		}

		num := strconv.Itoa(i)
		count := 0
		for _, w := range num {
			if w == '6' {
				count++
			} else {
				count = 0
			}

			if count == 3 {
				sss = append(sss, i)
				break
			}
		}
	}

	input := 0
	fmt.Fscanln(reader, &input)
	fmt.Fprintln(writer, sss[input-1])
}
