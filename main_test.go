package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	assert.Equal(t, 4, numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	assert.Equal(t, 6, numIdenticalPairs([]int{1, 1, 1, 1}))
	assert.Equal(t, 0, numIdenticalPairs([]int{1, 2, 3}))
}
