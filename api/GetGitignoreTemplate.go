package api

// Fetch the gitignore template from the GitHub API
func GetGitignoreTemplate(name string) (GitignoreTemplateResponse, error) {
	endpoint := "gitignore/templates/" + name
	return request[GitignoreTemplateResponse](endpoint)
}
