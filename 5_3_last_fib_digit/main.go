package main

import "fmt"

func main() {
	var n, m int64
	fmt.Scan(&n, &m)
	current := int64(1)
	prev := int64(0)

	if n == 1 {
		fmt.Println(1)
		return
	}
	i := n
	for n > 1 {
		prev, current = current, (prev+current)%m
		n--
		fmt.Printf("i: %d, mod: %d\n", i-n, current)
	}
	fmt.Println(current)
}
