package main

func findPermutationDifference(s string, t string) int {
	mapS := make(map[string]int)
	sum := 0
	for i, val := range s {
		mapS[string(val)] = i
	}

	for i, val := range t {
		curr, ok := mapS[string(val)]
		if ok {
			sum += Abs(curr - i)
		}
	}

	return sum
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
