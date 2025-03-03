package main

import "fmt"

func AskQuestion(question string) bool {
	var answer string

	fmt.Printf("%s ", question)
	fmt.Scanln(&answer)

	if answer == "Y" || answer == "y" {
		return true
	}

	return false
}
