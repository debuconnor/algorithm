package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type user struct {
	name string
	age int
	reg int
}

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int

	fmt.Fscanln(r, &n)

	users := make([]user, n)

	for i := 0; i < n; i++{
		fmt.Fscan(r, &users[i].age)
		fmt.Fscan(r, &users[i].name)
		users[i].reg = i
	}

	sort.SliceStable(users, func(i, j int) bool{
		return users[i].age < users[j].age
	})
	
	sort.SliceStable(users, func(i, j int) bool{
		if users[i].age == users[j].age{
			return users[i].reg < users[j].reg
		}
		return false
	})

	for _, v := range users{
		fmt.Fprintf(w, "%d %s\n", v.age, v.name)
	}
}
