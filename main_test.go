package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	assert.Equal(t, 6, maxFreqSum("successes"))
	assert.Equal(t, 3, maxFreqSum("aeiaeia"))
}
