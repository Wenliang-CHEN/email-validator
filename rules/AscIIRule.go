package rules

import ("regexp")

type AscIIRule struct {}

func (rule AscIIRule) GetPattern() string {
	return "^[[:ascii:]]*$"
}

func (rule AscIIRule) GetErrorMessage(sample string) string {
	return regexp.MustCompile(rule.GetPattern()).FindStringSubmatch(sample)[1]
}
