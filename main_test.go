package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	assert.Equal(t, 2, findPermutationDifference("abc", "bac"))
	assert.Equal(t, 12, findPermutationDifference("abcde", "edbac"))
}
