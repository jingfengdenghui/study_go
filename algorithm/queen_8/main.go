package main

import (
	"fmt"
)

func main() {
	queenList := queen8()
	a := make([][]int, 8, 8)
	for i := 0; i < 8; i++ {
		a[i] = make([]int, 8, 8)
		for j := 0; j < 8; j++ {
			if queenList[j] == i {
				a[i][j] = 1
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Print("\n")
	}

}
func queen8() []int {
	queenList := make([]int, 8, 8)
	for i := range queenList {
		queenList[i] = -1
	}
	queen(queenList, 0)
	return queenList
}

func queen(queenList []int, x int) bool {
	if x == len(queenList) {
		return true
	}
	for y := 0; y < len(queenList); y++ {
		canPut := true
		for i := 0; i < x; i++ {
			// 同行
			if queenList[i] == y {
				canPut = false
				break
			}
			//对角线
			if x-i == queenList[i]-y || x-i == y-queenList[i] {
				canPut = false
				break
			}
		}
		if canPut {
			queenList[x] = y
			if queen(queenList, x+1) {
				return true
			}
		}
	}
	return false
}
