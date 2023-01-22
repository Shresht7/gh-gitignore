package helpers

import "testing"

func TestContains(t *testing.T) {

	// Test slice of strings
	slice := []string{"Go", "Java", "JavaScript", "Python", "Ruby", "Rust", "Scala", "Swift"}

	// Test cases for Contains
	testCases := []struct {
		input    string
		expected bool
	}{
		{"Go", true},
		{"GO", false},
		{"Java", true},
		{"JavaScript", true},
		{"Python", true},
		{"Ruby", true},
		{"Rust", true},
		{"Scala", true},
		{"Swift", true},
		{"go", false},
		{"ja", false},
		{"pyt", false},
		{"ru", false},
		{"s", false},
		{"c", false},
		{"rust", false},
		{"gogo", false},
	}

	// Iterate over all test cases
	for _, testCase := range testCases {
		// Check if the expected and actual results are equal
		if Contains(slice, testCase.input) != testCase.expected {
			t.Errorf("Case: %s, Expected: %v, Actual: %v", testCase.input, testCase.expected, Contains(slice, testCase.input))
		}
	}

}
