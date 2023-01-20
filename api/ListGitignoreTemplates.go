package api

import "github.com/cli/go-gh"

// List all the gitignore templates from the GitHub API
func ListGitignoreTemplates() ([]string, error) {

	// Get the gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return nil, err
	}

	// gitignore templates endpoint
	endpoint := "gitignore/templates"

	// Get the list of gitignore templates
	gitignoreList := []string{}
	err = client.Get(endpoint, gitignoreList)
	if err != nil {
		return nil, err
	}

	return gitignoreList, nil

}
