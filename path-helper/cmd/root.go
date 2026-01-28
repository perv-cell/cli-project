/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pelp",
	Short: "The commands are:",
	Long:  `"pelp is a tool for work ENV, PATH"`,
}

func NewExecute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	/*rootCmd.SetHelpTemplate(`
	For help use "pelp -h or --help" for more information about that topic.

	Usage:
	pelp <command> [arguments]

	The commands are:

	  addpath   add our path in common PATH
		Flags:
			--temp	adds to PATH while the terminal is running
			--const	adds the PATH variable to your system configuration.
				It will remain in the PATH until you remove it.
			arguments-1: the path to add to PATH

		Example:
			pelp addpath --temp ~/mypackage


	  rmpath   remove our path in common PATH`,
		)*/
}
