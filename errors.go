package xssvalidator

import "errors"

var (
	ErrBracketRule       = errors.New("bracket rule triggered, input contains one of the following characters: (){}[]")
	ErrForbiddenKeywords = errors.New("forbidden keywords triggered, input contains one of the following keywords in a vulnerable format: alert, prompt")
)
