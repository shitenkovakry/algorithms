package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}

	expected := 3
	actual := FindTheMaximumNumberOfCheks(array)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	array := []int{}

	expected := 0
	actual := FindTheMaximumNumberOfCheks(array)

	assert.Equal(t, expected, actual)
}
