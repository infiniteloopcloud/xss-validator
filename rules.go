package xssvalidator

import (
	"html"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var keywords = []string{
	`(alert).*(\()`,
	`(prompt).*(\()`,
	`(eval).*(\()`,
	`(window).*(\[)`,
	`<script`,
	`</script`,
	`<x`,
	`<X`,
	`<http`,
	`(function).*(\()`,
	`<iframe`,
	`(href).*(=)`,
	`<br>`,
	"alert`",
	`(find).*(\()`,
	`(top).*(\[)`,
	`(vibrate).*(\()`,
	`<object`,
	`<embed`,
	`<img`,
	`<layer`,
	`<style`,
	`<meta`,
	`=".*"`,
	`<html`,
	`(echo).*(\()`,
	`(confirm).*(\()`,
	`(write).*(\()`,
	`</svg`,
	`<div`,
	`</image`,
	`form>`,
	`(vectors).*(\()`,
	`<body`,
	`(url).*(\()`,
	`math>`,
	`-->`,
	`<!--`,
	`<!attlist`,
	`<label`,
	`<%`,
	`xmp>`,
	`template>`,
	`<!doctype`,
	`=confirm`,
}

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
	re, err := regexp.Compile(strings.Join(keywords, "|"))
	if err != nil {
		return err
	}

	if re.MatchString(input) {
		return ErrForbiddenKeywords
	}
	return nil
}

type ForbiddenHTMLUnescapeStringKeywords struct{}

func (f ForbiddenHTMLUnescapeStringKeywords) Check(input string) error {
	re, err := regexp.Compile(strings.Join(keywords, "|"))
	if err != nil {
		return err
	}

	if re.MatchString(html.UnescapeString(input)) {
		return ErrForbiddenKeywords
	}
	return nil
}

type ForbiddenURLQueryUnescapeKeywords struct{}

func (f ForbiddenURLQueryUnescapeKeywords) Check(input string) error {
	re, err := regexp.Compile(strings.Join(keywords, "|"))
	if err != nil {
		return err
	}

	decoded, _ := url.QueryUnescape(input)
	if re.MatchString(decoded) {
		return ErrForbiddenKeywords
	}
	return nil
}

type ForbiddenUnicodeKeywords struct{}

func (f ForbiddenUnicodeKeywords) Check(input string) error {
	re, err := regexp.Compile(strings.Join(keywords, "|"))
	if err != nil {
		return err
	}

	in := `"` + input + `"`

	decoded, _ := strconv.Unquote(in)
	if re.MatchString(decoded) {
		return ErrForbiddenKeywords
	}
	return nil
}

type ForbiddenLowercaseKeywords struct{}

func (f ForbiddenLowercaseKeywords) Check(input string) error {
	re, err := regexp.Compile(strings.Join(keywords, "|"))
	if err != nil {
		return err
	}

	decoded := strings.ToLower(input)
	if re.MatchString(decoded) {
		return ErrForbiddenKeywords
	}
	return nil
}
