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

	var n, m int
	fmt.Fscan(r, &n, &m)

	sites := make(map[string]string)
	url := ""
	pass := ""
	result := make([]string, m)

	for i := 0; i < n; i++{
		fmt.Fscan(r, &url, &pass)
		sites[url] = pass
	}

	for i := 0; i < m; i++{
		fmt.Fscan(r, &url)
		result = append(result, sites[url])
	}

	for _, v := range result{
		if v != "" {
			fmt.Fprintln(w, v)
		}
	}
}