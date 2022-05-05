package main

import (
	"bufio"
	"fmt"
	"math"
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
	max := 0
	sum := 0
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &temp)
		lan = append(lan, temp)

		sum += temp
		if temp > max {
			max = temp
		}

		if temp < min {
			min = temp
		}
	}
	result := 0

	if k == 1{
		result = min / n
	} else {
		result = getLans(&lan, n, min, max, k, n)
	}

	fmt.Fprintln(writer, result)
}

func getLans(arr *[]int, goal int, min int, max int, k int, n int) int {
	m := 0
	for {
		count := 0

		for _, v := range *arr {
			count += v / powInt(2, m)
		}

		if count < goal{
			break
		}
		m++
	}
	return binSearch(powInt(2, m - 1), goal, arr, m)
}


func binSearch(l int, goal int, arr *[]int, depth int) int{
	count := 0

	if depth == 0{	
		for _, v := range *arr{
			count += v / (l + 1)
		}

		if goal <= count{
			return l + 1
		}

		return l
	}

	count = 0
	
	for _, v := range *arr{
		count += v / (l + powInt(2, depth))
	}

	if count < goal {
		l = binSearch(l, goal, arr, depth - 1)
	} else if count >= goal{
		l = binSearch(l + powInt(2, depth), goal, arr, depth - 1)
	}
	
	return l
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}