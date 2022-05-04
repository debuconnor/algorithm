package main

import (
	"bufio"
	"fmt"
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

	result := getLans(&lan, n, min, min, 0, false)

	fmt.Fprintln(writer, result)
}

func getLans(arr *[]int, goal int, length int, min int, count int, flag bool) int {
	if count != goal && flag {
		return length * 2
	}

	_count := 0

	for _, v := range *arr {
		_count += v / length
	}

	if _count < goal {
		length /= 2
	} else if _count > goal {
		length += (min + length) / 2
	} else if goal == _count {
		flag = true
		length++
	}

	return getLans(arr, goal, length, min, _count, flag)
}

func doubleCheck(arr *[]int, goal int, length int, count int) int {
	if count != goal {
		return length - 1
	}

	_count := 0
	_length := length + 1
	for _, v := range *arr {
		_count += v / _length
	}

	return doubleCheck(arr, goal, _length, _count)
}
