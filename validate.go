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

	rules := slice.Reject(GetRules(), os.Args[1], match)
	if len(rules) == 0 {
		fmt.Println("Email valid")
		os.Exit(0)
	}

	fmt.Println("Email invalid: ")
	slice.Each(rules, func(rule Rule) {
		fmt.Printf("%v\n", rule.ErrorMsg)
	})
}

var match = func(rule Rule, sample string) bool {
	match, err := regexp.MatchString(rule.Pattern, sample)

	if err != nil {
		return false
	}
	return match
}

