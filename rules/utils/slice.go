package slice

import (
	. "validator/rules"
)

func Reject(haystack []Rule, sample string, f func(rule Rule, sample string) bool) []Rule {
	target := make([]Rule, 0)
	for _, rule := range haystack {
		if f(rule, sample) == false {
			target = append(target, rule)
		}
	}
	return target
}

func Each(haystack []Rule, f func(rule Rule))  {
	for _, v := range haystack {
		f(v)
	}
}

