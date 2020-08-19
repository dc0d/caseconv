package caseconv_test

import (
	"fmt"

	"github.com/dc0d/caseconv"
)

func ExampleToCamel() {
	fmt.Println(caseconv.ToCamel("Test String v2"))

	// Output:
	// testStringV2
}

func ExampleToCamel_from_snake_case() {
	fmt.Println(caseconv.ToCamel("test_string_v2"))

	// Output:
	// testStringV2
}

func ExampleToCamel_from_upper_case() {
	fmt.Println(caseconv.ToCamel("TEST STRING V2"))

	// Output:
	// testStringV2
}

func ExampleToKebab() {
	fmt.Println(caseconv.ToKebab("Test String v2"))

	// Output:
	// test-string-v2
}

func ExampleToPascal() {
	fmt.Println(caseconv.ToPascal("test String v2"))

	// Output:
	// TestStringV2
}

func ExampleToSnake() {
	fmt.Println(caseconv.ToSnake("test String v2"))

	// Output:
	// test_string_v2
}
