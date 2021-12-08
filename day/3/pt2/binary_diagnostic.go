package main

import (
	"fmt"
	"io"
	"strconv"
)

func select_rows_with_position_byte(position_byte byte, position int, report []string) []string {
	var newReport []string
	for _, line := range report {
		if line[position] == position_byte {
			newReport = append(newReport, line)
		}
	}

	return newReport
}

func determine_most_common_position_byte(report []string, position int) byte {
	var zeros, ones = 0, 0
	var elem byte

	for _, line := range report {
		elem = line[position]
		if elem == '0' {
			zeros++
		} else {
			ones++
		}
		// fmt.Println(zeros, ones)

	}
	if zeros <= ones {
		return '1'
	} else {
		return '0'
	}
}

func find_oxygen_gen_rating(report_original []string) string {
	var report = report_original

	for i := 0; i < len(report[0]); i++ {
		if determine_most_common_position_byte(report, i) == '1' {
			report = select_rows_with_position_byte('1', i, report)
		} else {
			report = select_rows_with_position_byte('0', i, report)
		}
		if len(report) == 1 {
			return report[0]
		}
		// fmt.Println(report)
	}

	panic("invalid input")
}

func find_co2_scrub_rating(report_original []string) string {
	var report = report_original

	for i := 0; i < len(report[0]); i++ {
		if determine_most_common_position_byte(report, i) == '1' {
			report = select_rows_with_position_byte('0', i, report)
		} else {
			report = select_rows_with_position_byte('1', i, report)
		}
		if len(report) == 1 {
			return report[0]
		}
		// fmt.Println(report)
	}

	panic("invalid input")
}

func main() {
	var report []string
	var reportLine string

	for {
		_, err := fmt.Scanf("%s\n", &reportLine)

		if err == io.EOF {
			break
		}
		report = append(report, reportLine)
	}

	// fmt.Println(find_oxygen_gen_rating(report), find_co2_scrub_rating(report))
	oxygen_gen_rating_dec, _ := strconv.ParseInt(find_oxygen_gen_rating(report), 2, 64)
	co2_scrub_rating_dec, _ := strconv.ParseInt(find_co2_scrub_rating(report), 2, 64)

	fmt.Printf("%d\n", oxygen_gen_rating_dec*co2_scrub_rating_dec)
}
