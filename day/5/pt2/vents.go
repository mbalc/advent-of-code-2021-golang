package main

import (
	"fmt"
	"io"
)

const MAX_COORD = 1000

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var fields [MAX_COORD][MAX_COORD]int
	var fromX, fromY, toX, toY int
	var crossingsCount = 0

	for {
		_, err := fmt.Scanf("%d,%d -> %d,%d\n", &fromX, &fromY, &toX, &toY)
		// fmt.Println(fromX, fromY, toX, toY)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if fromX == toX {
			for i := min(fromY, toY); i <= max(fromY, toY); i++ {
				fields[i][fromX]++
			}
		} else if fromY == toY {
			for i := min(fromX, toX); i <= max(fromX, toX); i++ {
				fields[fromY][i]++
			}
		} else if fromX-toX == fromY-toY {
			var i, j = min(fromX, toX), min(fromY, toY)
			for i <= max(fromX, toX) {
				fields[j][i]++
				i++
				j++
			}
		} else if fromX-toX == -(fromY - toY) {
			var i, j = min(fromX, toX), max(fromY, toY)
			for i <= max(fromX, toX) {
				fields[j][i]++
				i++
				j--
			}
		} else {
			fmt.Println(fromX, fromY, "->", toX, toY)
			panic("bad input")
		}
	}

	// for _, row := range fields {
	// 	fmt.Println(row)
	// }

	for _, row := range fields {
		for _, coverage := range row {
			if coverage > 1 {
				crossingsCount++
			}
		}
	}

	fmt.Println(crossingsCount)
}
