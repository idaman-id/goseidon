package validation

type Data = map[string]interface{}
type Rule = map[string]interface{}

type ValidationRuleDto struct {
	Locale string
	Data   Data
	Rule   Rule
}

type ValidationStructDto struct {
	Locale string
	Struct interface{}
}
