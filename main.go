package main

import "math"

func findPermutationDifference(s string, t string) int {
	mapS := make(map[string]int)
	sum := 0
	for i, val := range s {
		mapS[string(val)] = i
	}

	for i, val := range t {
		curr, ok := mapS[string(val)]
		if ok {
			sum += int(math.Abs(float64(curr) - float64(i)))
		}
	}

	return sum
}
