package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputone_Returntwo(t *testing.T) {
	assert.Equal(t, 2, TwoSum(1))
}

func TestInputtwo_Returnthree(t *testing.T) {
	assert.Equal(t, 3, TwoSum(2))
}
