package helpers

// Returns a boolean indicating whether the provided item is in the provided slice
func Contains(slice []string, item string) bool {
	for _, listItem := range slice {
		if listItem == item {
			return true
		}
	}
	return false
}
