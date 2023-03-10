package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

//	============
//	ROOT COMMAND
//	============

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-gitignore",
	Short: "Generate .gitignore files for your projects",
	Long: `Generate .gitignore files for your projects using GitHub's gitignore template API.
	
Examples:
	gh gitignore list
	gh gitignore create Go
	gh gitignore view Node
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-gitignore.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
