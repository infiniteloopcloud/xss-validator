package xssvalidator

import "errors"

var (
	ErrBracketRule = errors.New("bracket rule triggered, input contains one of the following characters: (){}[]")
)
