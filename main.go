package main

func maxFreqSum(s string) int {
	Empty := struct{}{}
	vowels := map[string]struct{}{
		"a": Empty,
		"i": Empty,
		"u": Empty,
		"e": Empty,
		"o": Empty,
	}
	vowelCounter := make(map[string]int)
	consonantCounter := make(map[string]int)
	for _, letter := range s {
		_, isVowel := vowels[string(letter)]
		if isVowel {
			cur, ok := vowelCounter[string(letter)]
			if ok {
				vowelCounter[string(letter)] = cur + 1
			} else {
				vowelCounter[string(letter)] = 1
			}
		} else {
			cur, ok := consonantCounter[string(letter)]
			if ok {
				consonantCounter[string(letter)] = cur + 1
			} else {
				consonantCounter[string(letter)] = 1
			}
		}
	}

	return findMax(vowelCounter) + findMax(consonantCounter)
}

func findMax(mapped map[string]int) int {
	max := 0
	for _, val := range mapped {
		if val > max {
			max = val
		}
	}
	return max
}
