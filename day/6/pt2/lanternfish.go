package main

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_DAY = 256
const MAX_TIMER = 10
const TIMER_AFTER_RESET = 6
const TIMER_AFTER_SPAWN = 8

func main() {
	var buffer string
	fmt.Scanln(&buffer)
	var splitBuffer = strings.Split(buffer, ",")
	var timerFishCountBuckets [MAX_TIMER]int
	for _, numString := range splitBuffer {
		initialTimer, _ := strconv.Atoi(numString)
		timerFishCountBuckets[initialTimer]++
		// fmt.Println(initialTimer, timerFishCountBuckets)
	}
	// fmt.Println(timerFishCountBuckets)
	for i := 0; i < MAX_DAY; i++ {
		fishToBearChildCount := timerFishCountBuckets[0]
		for n := 1; n < MAX_TIMER; n++ {
			timerFishCountBuckets[n-1] = timerFishCountBuckets[n]
		}
		timerFishCountBuckets[TIMER_AFTER_RESET] += fishToBearChildCount
		timerFishCountBuckets[TIMER_AFTER_SPAWN] += fishToBearChildCount

		// fmt.Println(timerFishCountBuckets)
	}

	var totalFishCount = 0
	for _, fishWithTimerTimeCount := range timerFishCountBuckets {
		totalFishCount += fishWithTimerTimeCount
	}
	fmt.Println(totalFishCount)
}
