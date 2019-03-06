package rules

type MaxLengthRule struct {}

func (rule AtSignRule) GetPattern() string {
	return "-a"
}

func (rule AtSignRule) GetErrorMessage(sample string) string {
	return "Must contains exactly one @"
}
