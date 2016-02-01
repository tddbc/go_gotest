package tddbc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSay(t *testing.T) {
	actual := Say("Wow!")
	expected := "Wow! TDD BootCamp!!"

	if actual != expected {
		t.Errorf("actual=%s, expect=%s", actual, expected)
	}
}

func TestSay_testify(t *testing.T) {
	actual := Say("Hello!")
	assert.Equal(t, "Hello! TDD BootCamp!!", actual, "they should be equal")
}
