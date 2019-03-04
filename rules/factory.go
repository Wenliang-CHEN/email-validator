package rules

func GetRules() []Rule {
	rules := make([]Rule, 0)
	rules = append(rules, AtSignRule{})
	rules = append(rules, AscIIRule{})
	return rules
}
