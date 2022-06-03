package xssvalidator

import (
	"regexp"
	"strings"
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

type ForbiddenKeywords struct{}

func (f ForbiddenKeywords) Check(input string) error {
	re, err := regexp.Compile(f.createRegexpString())
	if err != nil {
		return err
	}

	if re.MatchString(input) {
		return ErrForbiddenKeywords
	}
	return nil
}

func (f ForbiddenKeywords) createRegexpString() string {
	keywords := []string{
		`(alert).*(\()`,
		`(prompt).*(\()`,
		`(eval).*(\()`,
		`(window).*(\[)`,
		`<script`,
		`<%78`,
		`<x`,
		`<X`,
		`<http`,
		`(function).*(\()`,
		`<iframe`,
		`(href).*(=)`,
		`<br>`, // TODO
	}
	return strings.Join(keywords, "|")
}
