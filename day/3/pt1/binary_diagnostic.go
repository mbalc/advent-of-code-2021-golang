package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var report []string
	var reportLine string
	var gamma, epsilon = "", ""

	for {
		_, err := fmt.Scanf("%s\n", &reportLine)

		if err == io.EOF {
			break
		}
		report = append(report, reportLine)
		// fmt.Println(report)
		// fmt.Println()
	}

	for len(report[0]) > 0 {
		var zeros, ones = 0, 0
		var elem byte
		for i, line := range report {
			elem, report[i] = line[0], line[1:]
			if elem == '0' {
				zeros++
			} else {
				ones++
			}
			// fmt.Println(zeros, ones)

		}
		if zeros < ones {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	gamma_decimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_decimal, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("%d\n", gamma_decimal*epsilon_decimal)
}
