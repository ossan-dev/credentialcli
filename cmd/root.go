/*
Copyright © 2024 Amrit Singh <amritsingh183@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

// FIXME: the project structure might be good.
/*
	I would have used the cobra-cli command to initialize it.
	Start by creating a go.mod with the command 'go mod init'.
	Then, issue the 'cobra-cli init'.
	I cannot see the root.go file.
	Then, keep adding commands & subcommands with the 'cobra-cli add <name of command>
*/
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "credential",
	Short:   "credential is a utility to generate credentials",
	Long:    "",
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetOutput(os.Stdout)

	rootCmd.AddCommand(passwordCmd)
	// Execute the Cobra command tree, parsing args and identifying the command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
