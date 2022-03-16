package tddbc

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestFizzBuzz(t *testing.T) {
	actual := FizzBuzz(1)
	assert.Equal(t, "1", actual, "this must return same number")
}

func TestFizzBuzz_3bai(t *testing.T) {
	actual := FizzBuzz(3)
	assert.Equal(t, "Fizz", actual, "3の倍数の場合はFizzを返す")
	actual = FizzBuzz(6)
	assert.Equal(t, "Fizz", actual, "3の倍数の場合はFizzを返す")
}
