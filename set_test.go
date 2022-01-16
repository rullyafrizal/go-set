package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var s1 Set
	assert.True(t, s1.add("A"))
	assert.False(t, s1.add("A"))
	assert.Equal(t, 1, s1.size())
}

func TestRemove(t *testing.T) {
	var s1 Set

	s1.add("A")
	s1.add("B")
	s1.add("C")
	s1.add("D")
	s1.add("E")

	assert.Equal(t, 5, s1.size())
	assert.True(t, s1.remove("A"))
	assert.Equal(t, 4, s1.size())
}
