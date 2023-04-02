package api

// List all the gitignore templates from the GitHub API
func ListGitignoreTemplates() ([]string, error) {

	// gitignore templates endpoint
	endpoint := "gitignore/templates"

	// Get the gitignore templates from the GitHub API
	response, err := request[[]string](endpoint)
	if err != nil {
		return response, err
	}

	// Return the gitignore templates
	return response, nil

}
