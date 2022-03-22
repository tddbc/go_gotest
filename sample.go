package tddbc

import "fmt"

func Say(greeting string) string {
	return greeting + " TDD BootCamp!!"
}

func FizzBuzz(value int) string {
	switch {
	case value%5 == 0:
		return "Buzz"
	case value%3 == 0:
		return "Fizz"
	default:
		return fmt.Sprintf("%d", value)
	}
}
