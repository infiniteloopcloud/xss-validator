package xssvalidator

import (
	"regexp"
)

type BracketRule struct{}

func (BracketRule) Check(input string) error {
	re, err := regexp.Compile(`\[.*?\]|{.*?}|\(.*?\)|<.*?>`)
	if err != nil {
		return err
	}

	if re.MatchString(input) {
		return ErrBracketRule
	}
	return nil
}
