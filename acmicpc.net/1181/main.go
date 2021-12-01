package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var tc int
	var s []string
	var smap map[string]bool

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &tc)

	smap = make(map[string]bool)
	temp := ""
	for i := 0; i < tc; i++ {
		fmt.Fscanln(reader, &temp)

		if _, ok := smap[temp]; !ok {
			smap[temp] = true
			s = append(s, temp)
		}
	}

	sort.Slice(s, func(i, j int) bool {
		if len(s[i]) == len(s[j]) {
			var ss []string
			ss = append(ss, s[i])
			ss = append(ss, s[j])
			sort.Strings(ss)
			return ss[0] == s[i]
		}
		return len(s[i]) < len(s[j])
	})

	for _, v := range s {
		fmt.Fprintln(writer, v)
	}
}
