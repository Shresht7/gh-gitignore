package api

import "github.com/cli/go-gh"

// Fetch the gitignore template from the GitHub API
func GetGitignoreTemplate(name string) (GitignoreTemplateResponse, error) {

	// Get the gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return GitignoreTemplateResponse{}, err
	}

	// gitignore templates endpoint
	endpoint := "gitignore/templates/" + name

	// Get the gitignore template
	gitignore := GitignoreTemplateResponse{}
	err = client.Get(endpoint, gitignore)
	if err != nil {
		return GitignoreTemplateResponse{}, err
	}

	return gitignore, nil

}
