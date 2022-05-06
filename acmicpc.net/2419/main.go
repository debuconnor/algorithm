package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscanln(r, &n, &m)

	position := make([]int, n)
	tmp := 0
	
	for i := range position{
		fmt.Fscanln(r, &tmp)
		position[i] = tmp
	}
	
	sua := 0 // x좌표
	candy := 0 // 먹은 개수
	
	fmt.Fprintln(w, getCandy(m, sua, candy, position))
}

func getCandy(m, sua, candy int, position []int) int{
	length := make([]int, len(position))
	
	for i := range length{
		length[i] = abs(position[i] - sua)
	}

	min := getMinIndex(length)
	if length[min] >= m{
		return candy
	}
	time.Sleep(time.Second / 2)
	fmt.Println(length, "min:",min, "  m:", m)
	return getCandy(m - length[min], position[min], candy + m - length[min], position)
}

func abs(num int) int{
	if num < 0{
		num -= num * 2
	}

	return num
}

func getMinIndex(arr []int) int{
	min := 999999999
	index := 0
	index = index

	for i, v := range arr{
		if v < min{
			min = v
			index = i
		}
	}

	return index
}