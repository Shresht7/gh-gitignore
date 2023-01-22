package helpers

import (
	"github.com/Shresht7/gh-gitignore/api"
)

// Check if the provided gitignore template name is valid
func IsValidTemplate(name string) bool {

	// Get list of all gitignore templates
	templates, err := api.ListGitignoreTemplates()
	if err != nil {
		return false
	}

	// Iterate over all provided gitignore templates and check if the provided name is valid
	for _, template := range templates {
		if name == template {
			return true
		}
	}

	// TODO: Suggest similar gitignore template names

	// If the provided name is not valid, return false
	return false

}
