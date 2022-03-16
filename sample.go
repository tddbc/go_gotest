package tddbc

import "fmt"

func Say(greeting string) string {
	return greeting + " TDD BootCamp!!"
}

func FizzBuzz(value int) string {
	return fmt.Sprintf("%d", value)
}
