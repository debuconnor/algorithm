package main

import (
	"fmt"
)

func main() {
	array := []int{1, 5, 2, 6, 3, 7, 4}
	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}
	var _return []int

	for _, cmd := range commands {
		if cmd[0] == cmd[1] {
			_return = append(_return, array[cmd[0]-1])
		} else {
			_array := make([]int, cmd[1]-cmd[0]+1)
			copy(_array, array[cmd[0]-1:cmd[1]])
			quickSort(_array, 0, len(_array)-1)
			_return = append(_return, _array[cmd[2]-1])
		}
	}

	fmt.Println(_return)
}

func partition(a []int, left, right int) int {
	p := a[right]
	for j := left; j < right; j++ {
		if a[j] < p {
			a[j], a[left] = a[left], a[j]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	return left
}

func quickSort(a []int, left, right int) {
	if left > right {
		return
	}

	p := partition(a, left, right)
	quickSort(a, left, p-1)
	quickSort(a, p+1, right)
}
