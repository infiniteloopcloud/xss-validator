# XSS Validator

[![Go Report Card](https://goreportcard.com/badge/github.com/infiniteloopcloud/xss-validator)](https://goreportcard.com/report/github.com/infiniteloopcloud/xss-validator) [![GoDoc](https://godoc.org/github.com/infiniteloopcloud/xss-validator?status.svg)](https://godoc.org/github.com/infiniteloopcloud/xss-validator) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

XSS Validator on input fields. It covers all the payloads from the `test/payloads.txt`.

### Usage

```shell
go get github.com/infiniteloopcloud/xss-validator
```

```go
package main

import xssvalidator "github.com/infiniteloopcloud/xss-validator"

func main() {
	err := xssvalidator.Validate("input_data", xssvalidator.DefaultRules...)
	if err != nil {
		// rule triggered
	}

	// or use selected
	err = xssvalidator.Validate("input_data", []xssvalidator.Rule{
		xssvalidator.ForbiddenKeywords{},
		xssvalidator.ForbiddenHTMLUnescapeStringKeywords{},
	}...)
	if err != nil {
		// rule triggered
	}
}
```

### Writing custom rules

Anything implements the `xssvalidator.Rule` can be a rule passed into the validator.

```go
package ownrule

import (
	"errors"
	"strings"

	xssvalidator "github.com/infiniteloopcloud/xss-validator"
)

var _ xssvalidator.Rule = AlertRule{}

type AlertRule struct{}

func (a AlertRule) Check(v string) error {
	if strings.Contains(v, "alert") {
		return errors.New("contains alert")
	}

	return nil
}
```
