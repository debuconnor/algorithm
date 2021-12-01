package main

import "fmt"

func main() {
	var w, h int
	var wbb, bwb [8]string
	initBoard(&wbb, &bwb)

	fmt.Scanf("%d %d\n", &h, &w)
	board := make([]string, h)

	for i := 0; i < h; i++ {
		fmt.Scanf("%s\n", &board[i])
	}

	min := 64
	for i := 0; i+8 <= h; i++ {
		for j := 0; j+8 <= w; j++ {
			wb := 64 - compare(board, wbb, j, i)
			bw := 64 - compare(board, bwb, j, i)

			if wb > bw {
				if bw < min {
					min = bw
				}
			} else if wb < bw {
				if wb < min {
					min = wb
				}
			}
		}
	}

	if min == 64 {
		min = 32
	}

	fmt.Println(min)
}

func initBoard(b1 *[8]string, b2 *[8]string) {
	for i, _ := range *b1 {
		if i%2 == 0 {
			b1[i] = "WBWBWBWB"
			b2[i] = "BWBWBWBW"
		} else {
			b1[i] = "BWBWBWBW"
			b2[i] = "WBWBWBWB"
		}
	}
}

func compare(target []string, dest [8]string, startW int, startH int) int {
	count := 0

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if target[i+startH][j+startW] == dest[i][j] {
				count++
			}
		}
	}

	return count
}
