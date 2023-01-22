package helpers

// Returns a boolean indicating whether the provided item is in the given slice
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
