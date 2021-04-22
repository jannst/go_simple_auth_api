package store

import "testing"

type TokenValidationTestData struct {
	input string
	shouldBeValid bool
}

var (
	testData = []TokenValidationTestData{
		{input: "", shouldBeValid: false},
		{input: "11112222333344445555666677778888", shouldBeValid: true},
		{input: "0123456789abcdef0123456789abcdef", shouldBeValid: true},
		{input: "0123456789abcdef01234_6789abcdef", shouldBeValid: false},
		{input: "11112222333344445555666677778888a", shouldBeValid: false},
		{input: "1111222233334444555566667777888", shouldBeValid: false},
		{input: "1111222233334444555566667777888zz", shouldBeValid: false},
	}
)

func TestValidateTokenFormat(t *testing.T) {
	for _, test := range testData {
		if (ValidateTokenFormat(test.input) == nil) != test.shouldBeValid {
			t.Errorf("expecting result to be %t on input %s", test.shouldBeValid, test.input)
		}
	}
}
