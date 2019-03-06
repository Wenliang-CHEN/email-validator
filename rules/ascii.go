package rules

import (
	"regexp"
	"strings"
)

type AscIIRule struct {}

func (rule AscIIRule) GetPattern() string {
	return "^[[:ascii:]]*$"
}

func (rule AscIIRule) GetErrorMessage(sample string) string {
	nonAscIIs := regexp.MustCompile(rule.getNegatePattern()).FindAllString(sample, -1)
	return "Char \""+strings.Join(nonAscIIs, ",")+"\"  not allowed"
}

func (rule AscIIRule) getNegatePattern() string {
	return "[^[:ascii:]]"
}

