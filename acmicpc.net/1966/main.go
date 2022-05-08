package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	const PRINT = -1
	var tc, doc, target int
	var importance, result []int

	fmt.Fscanln(r, &tc)

	for i := 0; i < tc; i++{
		fmt.Fscan(r, &doc, &target)
		importance = make([]int, doc)

		for j := 0; j < doc; j++{
			fmt.Fscan(r, &importance[j])
		}

		printCount := 0
		delVal := 0

		// fmt.Println(">>Round: ", i, " / doc qty: ", doc, " / target Position: ", target, importance)

		for target != PRINT {
			maxIndex := getMaxIndex(importance)
			
			for i := maxIndex; i >= 0; i--{
				importance, delVal = delete(importance)

				if target == PRINT{
					target = (doc - 1) - printCount
				}

				if i == 0{
					printCount++
					target--
					break
				} else {
					importance = insert(importance, delVal)
				}

				target--
			}
		}

		result = append(result, printCount)
	}

	for _, v := range result{
		fmt.Fprintln(w, v)
	}
}

func insert(arr []int, val int) []int{
	arr = append(arr, val)

	return arr
}

func delete(arr []int) ([]int, int){
	oldVal := arr[0]
	arr = arr[1:]

	return arr, oldVal
}

func getMaxIndex(arr []int) int{
	index := 0
	max := 0

	for i, v := range arr{
		if v > max && max != v{
			max = v
			index = i
		}
	}

	return index
}