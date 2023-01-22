package helpers

import "fmt"

// Handle invalid gitignore template name case
func HandleInvalidTemplate(name string, templates []string) {
	fmt.Printf("Invalid gitignore template name: %s\n", name)

	// Get suggestions for the given name
	suggestions := GetSuggestions(name, templates)

	// Print suggestions, if any
	if len(suggestions) > 0 {
		fmt.Println("\nDid you mean:")
		for _, suggestion := range suggestions {
			fmt.Println("\t" + suggestion)
		}
	}

	fmt.Println("\nUse the \"list\" command to view all available gitignore templates.")
}
