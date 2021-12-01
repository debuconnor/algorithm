package main

import "fmt"

func main() {
	var x, y, w, h int
	result := 9999
	fmt.Scanf("%d %d %d %d", &x, &y, &w, &h)

	if w-x < result {
		result = w - x
	}
	if x < result {
		result = x
	}
	if h-y < result {
		result = h - y
	}
	if y < result {
		result = y
	}

	fmt.Println(result)
}
