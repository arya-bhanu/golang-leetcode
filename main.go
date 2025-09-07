package main

func numJewelsInStones(jewels string, stones string) int {
	mapJewels := make(map[string]struct{})
	totalCounter := 0
	for _, jewel := range jewels {
		mapJewels[string(jewel)] = struct{}{}
	}

	for _, stone := range stones {
		_, ok := mapJewels[string(stone)]
		if ok {
			totalCounter++
		}
	}

	return totalCounter
}
