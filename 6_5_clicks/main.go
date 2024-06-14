package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	clicks := make([]int, n)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&clicks[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&cost[i])
	}
	sort.Ints(clicks)
	sort.Ints(cost)
	var res int64
	for i := 0; i < n; i++ {
		res += int64(clicks[i] * cost[i])
	}
	fmt.Println(res)
}
