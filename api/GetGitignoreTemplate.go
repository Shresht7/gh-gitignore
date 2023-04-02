package api

// Fetch the gitignore template from the GitHub API
func GetGitignoreTemplate(name string) (GitignoreTemplateResponse, error) {

	// gitignore templates endpoint
	endpoint := "gitignore/templates/" + name

	// Get the gitignore template from the GitHub API
	response, err := request[GitignoreTemplateResponse](endpoint)
	if err != nil {
		return GitignoreTemplateResponse{}, err
	}

	// Return the gitignore template
	return response, nil

}
