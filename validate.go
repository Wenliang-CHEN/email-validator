package main

import (
	. "validator/rules"
	slice "validator/rules/utils"
	"fmt"
	"regexp"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter an email address")
		os.Exit(0)
	}

	sample := os.Args[1]
	rules := slice.Reject(GetRules(), sample, match)
	if len(rules) == 0 {
		fmt.Println("Email valid")
		os.Exit(0)
	}

	fmt.Println("Email invalid: ")
	slice.Each(rules, func(rule Rule) {
		fmt.Printf("%v\n", rule.GetErrorMessage(sample))
	})
}

var match = func(rule Rule, sample string) bool {
	match, err := regexp.MatchString(rule.GetPattern(), sample)

	if err != nil {
		return false
	}
	return match
}

