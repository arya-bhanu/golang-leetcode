package main

func numIdenticalPairs(nums []int) int {
	counter := 0
	for i, val := range nums {
		j := i + 1
		for j < len(nums) {
			if val == nums[j] {
				counter++
			}
			j++
		}
	}
	return counter
}
