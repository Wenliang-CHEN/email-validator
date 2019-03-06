package rules

type AtSignRule struct {}

func (rule AtSignRule) GetPattern() string {
	return "^[^@]*@[^@]*$"
}

func (rule AtSignRule) GetErrorMessage(sample string) string {
	return "Must contains exactly one @"
}
