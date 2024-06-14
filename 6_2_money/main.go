package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n/10 + (n%10)/5 + (n % 5))
}
