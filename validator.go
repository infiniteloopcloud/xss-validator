package xssvalidator

var DefaultRules []Rule = []Rule{
	//BracketRule{},
	ForbiddenKeywords{},
}

type Rule interface {
	Check(string) error
}

func Validate(input string, rules ...Rule) error {
	for _, rule := range rules {
		if err := rule.Check(input); err != nil {
			return err
		}
	}
	return nil
}
