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
	var actual string
	table := []struct {
		description string
		input       int
		output      string
	}{
		{
			"数字をそのまま返す",
			1,
			"1",
		},
		{
			"3の倍数の場合はFizzを返す",
			3,
			"Fizz",
		},
		{
			"3の倍数の場合はFizzを返す",
			6,
			"Fizz",
		},
		{
			"3の倍数の場合はFizzを返す",
			111111,
			"Fizz",
		},
	}
	for _, v := range table {
		actual = FizzBuzz(v.input)
		assert.Equal(t, v.output, actual, v.description)
	}
}
