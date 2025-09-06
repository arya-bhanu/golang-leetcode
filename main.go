package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := []int{3, 2, 1, 10, 12, 200, 100}
	sort.Ints(tests)
	fmt.Println(tests)
}

func getSneakyNumbers(nums []int) []int {
	var duplicate []int
	sort.Ints(nums)
	for i, val := range nums {
		if i+1 < len(nums) {
			if val == nums[i+1] {
				duplicate = append(duplicate, val)
			}
		}
	}
	return duplicate
}
