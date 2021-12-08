package main

import (
	"fmt"
	"io"
)

const DEPTH_NOT_INITIALIZED = -42

func main() {
	var last_depth = DEPTH_NOT_INITIALIZED
	var curr_depth, depth_increase_count = 0, 0

	for {
		_, err := fmt.Scanf("%d", &curr_depth)

		if err == io.EOF {
			break
		}
		if last_depth != DEPTH_NOT_INITIALIZED && last_depth < curr_depth {
			depth_increase_count++
		}

		last_depth = curr_depth
	}

	fmt.Printf("%d\n", depth_increase_count)
}
