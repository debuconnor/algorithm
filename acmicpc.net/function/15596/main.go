package main

func main() {
	var a [3]int
	result := 0

	for _, v := range a {
		result += v
	}

	return result
}
