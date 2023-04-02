package helpers

import (
	"testing"
)

func TestListExamples(t *testing.T) {

	examples := []string{
		"gh gitignore list",
		"gh gitignore view Node",
		"gh gitignore create Rust",
	}
	result := ListExamples(examples)

	// ! Mind the 2 spaces before each example. This output is what looks good with cobra.
	expected := `  gh gitignore list
  gh gitignore view Node
  gh gitignore create Rust
`

	if result != expected {
		t.Errorf("want:\n%s\n\ngot:\n%s", expected, result)
	}
}
