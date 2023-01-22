package helpers

import "strings"

// Get suggestions for the given input string from the given slice
func GetSuggestions(input string, slice []string) []string {
	var suggestions []string
	for _, item := range slice {
		if strings.Contains(strings.ToLower(item), strings.ToLower(input)) {
			suggestions = append(suggestions, item)
		}
	}
	return suggestions
}
