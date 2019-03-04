package rules

type Rule interface {
	GetPattern() string
	GetErrorMessage(sample string) string
}
