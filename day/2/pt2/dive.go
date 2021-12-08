package main

import (
	"fmt"
	"io"
)

func main() {
	var operation_type string
	var amount int
	var curr_depth, curr_x_pos, aim = 0, 0, 0

	for {
		_, err := fmt.Scanf("%s %d\n", &operation_type, &amount)

		if err == io.EOF {
			break
		}

		if operation_type == "forward" {
			curr_x_pos += amount
			curr_depth += (amount * aim)
		}
		if operation_type == "up" {
			aim -= amount
		}
		if operation_type == "down" {
			aim += amount
		}
	}

	fmt.Println(curr_depth * curr_x_pos)
}
