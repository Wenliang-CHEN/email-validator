package rules

func GetRules() []Rule {
	rules := make([]Rule, 0)
	rules = append(rules, Rule{"^[^@]*@[^@]*$", "You need 1 and only 1 @ sign"})
	rules = append(rules, Rule{"^[[:ascii:]]*$", "Please use ASCII characters only"})
	return rules
}
