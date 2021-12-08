package main

import (
	"fmt"
	"io"
)

func main() {
	var operation_type string
	var amount int
	var curr_depth, curr_x_pos = 0, 0

	for {
		_, err := fmt.Scanf("%s %d\n", &operation_type, &amount)

		fmt.Println(operation_type, amount)
		if err == io.EOF {
			break
		}

		if operation_type == "forward" {
			curr_x_pos += amount
		}
		if operation_type == "up" {
			curr_depth -= amount
		}
		if operation_type == "down" {
			curr_depth += amount
		}
		fmt.Println(curr_depth, curr_x_pos)
	}

	fmt.Println(curr_depth * curr_x_pos)
}
