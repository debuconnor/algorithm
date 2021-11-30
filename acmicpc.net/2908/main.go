package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	var aa, bb []byte

	fmt.Scanf("%s %s", &a, &b)

	aa = []byte(a)
	bb = []byte(b)

	aa[0], aa[2] = aa[2], aa[0]
	bb[0], bb[2] = bb[2], bb[0]
	x, _ := strconv.Atoi(string(aa))
	y, _ := strconv.Atoi(string(bb))

	if x > y {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}
}
