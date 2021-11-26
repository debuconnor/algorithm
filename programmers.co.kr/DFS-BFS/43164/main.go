package main

import (
	"fmt"
	"sort"
)

func main() {
	var tickets = [][]string{{"ICN", "SFO"}, {"ICN", "ATL"}, {"SFO", "ATL"}, {"ATL", "ICN"}, {"ATL", "SFO"}}
	var result []string
	result = append(result, "ICN")

	fmt.Println(getDestination(tickets, result[0]))
}

func getDestination(tickets [][]string, departure string) string {
	var country []string

	for _, v := range tickets {
		if v[0] == departure {
			country = append(country, v[1])
		}
	}

	sort.Strings(country)

	return country[0]
}

func useTicket(tickets *[][]string, used *[]string) {

}
