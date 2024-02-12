package api

// List all the gitignore templates from the GitHub API
func ListGitignoreTemplates() ([]string, error) {
	endpoint := "gitignore/templates"
	return request[[]string](endpoint)
}
