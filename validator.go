package xssvalidator

var rules []Rule = []Rule{
	BracketRule{},
}

type Rule interface {
	Check(string) error
}

func Validate(input string, additionalRules ...Rule) error {
	for _, rule := range rules {
		if err := rule.Check(input); err != nil {
			return err
		}
	}
	return nil
}
