package helpers

import "testing"

func TestGetSuggestions(t *testing.T) {
	// Test slice of strings
	slice := []string{"Go", "Java", "JavaScript", "Python", "Ruby", "Rust", "Scala", "Swift"}

	// Test cases for GetSuggestions
	testCases := []struct {
		input    string
		expected []string
	}{
		{"go", []string{"Go"}},
		{"GO", []string{"Go"}},
		{"ja", []string{"Java", "JavaScript"}},
		{"PYT", []string{"Python"}},
		{"ru", []string{"Ruby", "Rust"}},
		{"s", []string{"JavaScript", "Rust", "Scala", "Swift"}},
		{"c", []string{"JavaScript", "Scala"}},
		{"Rust", []string{"Rust"}},
		{"GOGO", []string{}},
	}

	// Iterate over all test cases
	for _, testCase := range testCases {
		// Get suggestions for the given input string from the given slice
		suggestions := GetSuggestions(testCase.input, slice)

		// Check if the expected and actual suggestions are equal
		if !equal(suggestions, testCase.expected) {
			t.Errorf("Expected: %v, Actual: %v", testCase.expected, suggestions)
		}
	}
}

// Check if two slices of strings are equal
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
