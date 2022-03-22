package tddbc

import "fmt"

func Say(greeting string) string {
	return greeting + " TDD BootCamp!!"
}

func FizzBuzz(value int) string {
	if value%5 == 0 {
		return "Buzz"
	}
	if value%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", value)
}
