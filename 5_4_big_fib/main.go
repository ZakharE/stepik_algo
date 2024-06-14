package main

import "fmt"

func main() {
	var n, m int64
	fmt.Scan(&n, &m)
	period := pisanoPeriod(m)
	n = n % int64(period)
	fmt.Println(lastDigit(n, m))
}

func lastDigit(n, m int64) int64 {
	current := int64(1)
	prev := int64(0)

	if n <= 1 {
		return n
	}
	for n > 1 {
		prev, current = current, (prev+current)%m
		n--
	}
	return current

}

func pisanoPeriod(m int64) int {
	current := int64(0)
	next := int64(1)
	period := 0
	for {
		nextOld := next
		next = (next + current) % m
		current = nextOld

		period++
		if current == 0 && next == 1 {
			break
		}
	}
	return period

}
