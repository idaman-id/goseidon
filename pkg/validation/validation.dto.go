package validation

type Data = map[string]interface{}
type Rule = map[string]interface{}

type ValidationRuleParam struct {
	Locale string
	Data   Data
	Rule   Rule
}

type ValidationStructParam struct {
	Locale string
	Struct interface{}
}
