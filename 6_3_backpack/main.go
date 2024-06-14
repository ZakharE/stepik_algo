package main

import (
	"fmt"
	"sort"
)

type good struct {
	Cost   int
	Weight int
}

type ByValue []good

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Cost*a[j].Weight > a[j].Cost*a[i].Weight }

func main() {
	var n int
	var capacity int64
	fmt.Scan(&n, &capacity)
	i := 0

	weights := make([]good, 0, n)
	for i < n {
		var g good
		fmt.Scan(&g.Cost, &g.Weight)
		weights = append(weights, g)
		i++
	}
	sort.Sort(ByValue(weights))
	fmt.Printf("weights: %v\n", weights)
	result := float64(0)
	for _, w := range weights {
		toPut := min(capacity, int64(w.Weight))
		fmt.Printf("toPut: %v\n", toPut)
		capacity -= toPut
		result += float64(toPut) * (float64(w.Cost) / float64(w.Weight))
		if capacity == 0 {
			break
		}
	}
	fmt.Println(result)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
