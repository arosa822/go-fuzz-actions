package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdderFunc(t *testing.T) {
	input := 10
	expected := 11
	result := adder(input)
	assert.Equal(t, expected, result)

}
