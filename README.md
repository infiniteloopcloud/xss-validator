# XSS Validator

XSS Validator on input fields. It covers all the payloads from the `test/payloads.txt`.

### Usage

```shell
go get github.com/infiniteloopcloud/xss-validator
```

```go
package main

import (
	"github.com/infiniteloopcloud/xss-validator"
)

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

	"github.com/infiniteloopcloud/xss-validator"
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
