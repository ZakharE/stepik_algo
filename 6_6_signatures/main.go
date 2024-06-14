package main

import (
	"fmt"
	"sort"
)

type segment struct {
	left  int
	right int
}
type ByRightEnd []segment

func (a ByRightEnd) Len() int           { return len(a) }
func (a ByRightEnd) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRightEnd) Less(i, j int) bool { return a[i].right < a[j].right }

func main() {
	var n int
	fmt.Scan(&n)
	segments := make([]segment, n)
	for i := range segments {
		var s segment
		fmt.Scan(&s.left, &s.right)
		segments[i] = s
	}
	p := points(segments)
	fmt.Println(len(p))
	for _, v := range p {
		fmt.Println(v)

	}

}

func points(segments []segment) []int {
	if len(segments) == 1 {
		return []int{segments[0].right}
	}

	sort.Sort(ByRightEnd(segments))
	points := make([]int, 0, len(segments))
	mostLeft := segments[0]
	for i := 1; i < len(segments); i++ {
		if mostLeft.right >= segments[i].left {
			continue
		}
		if mostLeft.right < segments[i].left {
			points = append(points, mostLeft.right)
			mostLeft = segments[i]

		}
	}
	points = append(points, mostLeft.right)
	return points
}
