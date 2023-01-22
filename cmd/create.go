package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-gitignore/api"
)

//	==============
//	CREATE COMMAND
//	==============

// Flags
var (
	// Destination of the gitignore file
	dest string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "init", "clone"},
	Short:   "Create a gitignore file",
	Long:    `Create a gitignore file for your project`,
	Run: func(cmd *cobra.Command, args []string) {

		// Contents of the gitignore file
		var contents []byte

		// If the gitignore file already exists, read its contents
		contents, _ = os.ReadFile(dest)

		// Iterate over the given arguments and append the contents of the gitignore files
		for _, name := range args {

			//	Get the gitignore template
			gitignore, err := api.GetGitignoreTemplate(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			//	Append the gitignore template to the contents
			contents = append(contents, []byte("\n"+gitignore.Source+"\n")...)

		}

		// Write to file
		err := os.WriteFile(dest, contents, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	//	Add create command
	rootCmd.AddCommand(createCmd)

	//	Add flags
	createCmd.Flags().StringVarP(&dest, "dest", "d", ".gitignore", "Destination of the gitignore file")
}
