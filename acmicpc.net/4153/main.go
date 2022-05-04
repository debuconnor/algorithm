package main

import "fmt"

func main() {
	var x, y, z int
	var r []string

	for {
		fmt.Scanf("%d %d %d", &x, &y, &z)

		if z < x{
			x, z = z, x
		}
		if z < y{
			y, z = z, y
		}
		
		if x == 0 && y == 0 && z == 0{
			break
		}

		if x * x + y * y == z * z{
			r = append(r, "right")
		} else{
			r = append(r, "wrong")
		}
	}

	for _, v := range r{
		fmt.Println(v)
	}
}