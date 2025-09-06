package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	assert.ElementsMatch(t, []int{0, 1}, getSneakyNumbers([]int{0, 1, 1, 0}))
	assert.ElementsMatch(t, []int{2, 3}, getSneakyNumbers([]int{0, 3, 2, 1, 3, 2}))
	assert.ElementsMatch(t, []int{4, 5}, getSneakyNumbers([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}
