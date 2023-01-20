package api

type GitignoreTemplateResponse struct {
	// Name of the gitignore template
	Name string `json:"name"`
	// Source of the gitignore template
	Source string `json:"source"`
}
