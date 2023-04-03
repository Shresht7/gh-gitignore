package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-gitignore/api"
	"github.com/Shresht7/gh-gitignore/helpers"
)

//	==============
//	CREATE COMMAND
//	==============

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add", "init", "clone"},
	Short:   "Create a gitignore file",
	Long:    `Create a gitignore file for your project`,
	Example: helpers.ListExamples([]string{
		"gh gitignore create Go",
		"gh gitignore init Go Python",
		"gh gitignore create Go Python -d .gitignore",
	}),
	Run: func(cmd *cobra.Command, args []string) {

		// Contents of the gitignore file
		var contents []byte

		// Get Flags
		dest, _ := cmd.Flags().GetString("dest")
		overwrite, _ := cmd.Flags().GetBool("overwrite")

		// If the gitignore file already exists, read its contents
		if !overwrite {
			contents, _ = os.ReadFile(dest)
		}

		//	Get the gitignore templates
		templates, err := api.ListGitignoreTemplates()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Iterate over the given arguments and append the contents of the gitignore files
		for _, name := range args {

			// Check if the provided gitignore template name is invalid
			if !helpers.Contains(templates, name) {
				helpers.HandleInvalidTemplate(name, templates)
				continue // Skip the iteration
			}

			//	Get the gitignore template
			gitignore, err := api.GetGitignoreTemplate(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			//	Append the gitignore template to the contents
			if len(contents) > 0 { // If the contents are not empty, append a new line
				contents = append(contents, []byte("\n")...)
			}
			contents = append(contents, []byte(gitignore.Source+"\n")...) // Append the gitignore template

		}

		// Write to file
		err = os.WriteFile(dest, contents, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	//	Add create command
	RootCmd.AddCommand(createCmd)

	//	Add flags
	createCmd.Flags().StringP("dest", "d", ".gitignore", "Destination of the gitignore file")
	createCmd.Flags().BoolP("overwrite", "o", false, "Overwrite the gitignore file")
}
