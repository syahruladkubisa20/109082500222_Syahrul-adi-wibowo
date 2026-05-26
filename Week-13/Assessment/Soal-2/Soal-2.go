package main

import "fmt"

func selectionSortNames(names *[100]string, goals *[100]int, assists *[100]int, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if goals[j] > goals[maxIdx] {
				maxIdx = j
			} else if goals[j] == goals[maxIdx] {
				if assists[j] > assists[maxIdx] {
					maxIdx = j
				}
			}
		}
		if maxIdx != i {
			names[i], names[maxIdx] = names[maxIdx], names[i]
			goals[i], goals[maxIdx] = goals[maxIdx], goals[i]
			assists[i], assists[maxIdx] = assists[maxIdx], assists[i]
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	var names [100]string
	var goals [100]int
	var assists [100]int
	for i := 0; i < n; i++ {
		fmt.Scan(&names[i], &goals[i], &assists[i])
	}
	selectionSortNames(&names, &goals, &assists, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s %d %d\n", i+1, names[i], goals[i], assists[i])
	}
}
