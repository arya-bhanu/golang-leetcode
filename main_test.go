package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	funcResult := SolveProblem()
	assert.Equal(t, "Hello World", funcResult)
}
