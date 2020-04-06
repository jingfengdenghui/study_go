package main

import (
	"fmt"
)

var b = 1

func main() {
	a := []int{1, 2, 5, 3, 7, 4}
	fmt.Print("a:", quick(a))
}

func quick(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	quickSort(a, 0, len(a)-1)
	return a
}

func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}
	nowStart := start
	nowEnd := end
	for nowStart < nowEnd {
		for nowStart < nowEnd {
			if a[nowStart] > a[nowEnd] {
				swap(a, nowStart, nowEnd)
				break
			} else {
				nowStart++
			}
		}
		for nowStart < nowEnd {
			if a[nowStart] > a[nowEnd] {
				swap(a, nowStart, nowEnd)
				break
			} else {
				nowEnd--
			}
		}
	}
	quickSort(a, start, nowStart-1)
	quickSort(a, nowStart+1, end)
}

func swap(a []int, i, j int) {
	k := a[i]
	a[i] = a[j]
	a[j] = k
}
