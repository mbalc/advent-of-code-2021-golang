package main

import (
	"fmt"
	"io"
)

func main() {
	var measurement_window [4]int
	var depth_increase_count = 0

	fmt.Scan(&measurement_window[0])
	fmt.Scan(&measurement_window[1])
	fmt.Scan(&measurement_window[2])

	for {
		_, err := fmt.Scan(&measurement_window[3])

		if err == io.EOF {
			break
		}
		var last_sum = measurement_window[0] + measurement_window[1] + measurement_window[2]
		var curr_sum = measurement_window[1] + measurement_window[2] + measurement_window[3]
		// fmt.Printf("cmp %d, %d [%d %d %d %d]\n", last_sum, curr_sum, measurement_window[0], measurement_window[1], measurement_window[2], measurement_window[3])

		if last_sum < curr_sum {
			depth_increase_count++
		}
		measurement_window[0], measurement_window[1], measurement_window[2] = measurement_window[1], measurement_window[2], measurement_window[3]
	}

	fmt.Printf("%d\n", depth_increase_count)
}
