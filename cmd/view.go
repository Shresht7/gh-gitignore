package cmd

import (
	"fmt"

	"github.com/Shresht7/gh-gitignore/api"
	"github.com/spf13/cobra"
)

//	============
//	VIEW COMMAND
//	============

var viewCmd = &cobra.Command{
	Use:     "view",
	Short:   "View the gitignore file",
	Long:    `View the gitignore file`,
	Args:    cobra.ExactArgs(1),
	Example: `gh gitignore view node`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get the gitignore template name
		name := args[0]

		//	Get the gitignore template
		gitignore, err := api.GetGitignoreTemplate(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the gitignore template
		fmt.Println(gitignore.Source)

	},
}

func init() {
	//	Add view command
	rootCmd.AddCommand(viewCmd)
}
