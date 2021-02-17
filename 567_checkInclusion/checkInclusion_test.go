package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_case01(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	assert.Equal(t, true, CheckInclusion(s1, s2))
}
func Test_case02(t *testing.T) {
	s1 := "ab"
	s2 := "eidboaoo"
	assert.Equal(t, false, CheckInclusion(s1, s2))
}
