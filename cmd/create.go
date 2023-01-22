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

		//	Get the gitignore template name
		name := args[0]

		//	Get the gitignore template
		gitignore, err := api.GetGitignoreTemplate(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Write license file
		os.WriteFile(dest, []byte(gitignore.Source), 0644)

	},
}

func init() {
	//	Add create command
	rootCmd.AddCommand(createCmd)

	//	Add flags
	createCmd.Flags().StringVarP(&dest, "dest", "d", ".gitignore", "Destination of the gitignore file")
}
