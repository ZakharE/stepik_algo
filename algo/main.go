package main 

import "fmt"
func main() {
	var a []uint64 
	a = scanSlice()
	var max, maxN  uint64

	for _, num := range a {
		if num >= max {
			maxN = max
			max = num
		} else if num > maxN {
			maxN = num
		}
	}

	fmt.Print(max* maxN)
}

func scanSlice() []uint64 {
	var n int
	fmt.Scan(&n)
	res := make([]uint64, n)
	for i:= 0; i < n; i++ {
		fmt.Scan(&res[i])
	}
	return res
}

