[![PkgGoDev](https://pkg.go.dev/badge/dc0d/caseconv)](https://pkg.go.dev/github.com/dc0d/caseconv) [![Go Report Card](https://goreportcard.com/badge/github.com/dc0d/caseconv)](https://goreportcard.com/report/github.com/dc0d/caseconv) [![Maintainability](https://api.codeclimate.com/v1/badges/a0306f7932dae43bbda5/maintainability)](https://codeclimate.com/github/dc0d/caseconv/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/a0306f7932dae43bbda5/test_coverage)](https://codeclimate.com/github/dc0d/caseconv/test_coverage)

# caseconv

This is a Go Module for snake, kebab, camel, pascal case conversion.

It can be used like:

```go
package main

import (
	"fmt"

	"github.com/dc0d/caseconv"
)

func main() {
	input := "The quick brown fox jumps over the lazy dog"

	fmt.Println(caseconv.ToCamel(input))
	fmt.Println(caseconv.ToPascal(input))
	fmt.Println(caseconv.ToKebab(input))
	fmt.Println(caseconv.ToSnake(input))
}
```

And the output would be:

```
theQuickBrownFoxJumpsOverTheLazyDog
TheQuickBrownFoxJumpsOverTheLazyDog
the-quick-brown-fox-jumps-over-the-lazy-dog
the_quick_brown_fox_jumps_over_the_lazy_dog
```

> Most of test cases are from [change-case](https://github.com/blakeembrey/change-case) node package - so far. But the goal was not to follow same conventions.