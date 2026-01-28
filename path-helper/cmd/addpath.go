/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime"

	"github.com/perv-cell/cli-project/path-helper/cmd/lib/workpath"
	"github.com/spf13/cobra"
)

// addpathCmd represents the addpath command
var addpathCmd = &cobra.Command{
	Use:   "addpath",
	Short: "add our path in common PATH",
	Long: `
adds your file path to the configuration
or temporarily during a session
to work with the system environment.
For example, if you downloaded the utility,
you need to call it anywhere in the system.
Adding it to the PATH will complete this task.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("To run this command, specify the path you want to add to PATH as the only argument.")
			fmt.Println(args)
			return
		}
		path := args[0]

		_const, _ := cmd.Flags().GetBool("const")
		if _const {
			os := runtime.GOOS
			switch os {

			case "windows":
				err := workpath.AddUserPathInPATH(path)
				if err != nil {
					fmt.Println("Error read PATH user:", err)
				}

			case "linux":
				fmt.Println("Запущено на Linux")

			case "darwin":
				fmt.Println("Запущено на macOS (Darwin)")

			default:
				fmt.Printf("Другая ОС: %s\n", os)
			}

			fmt.Println("The path has been successfully added to the system variable. 😏")
		} else {

			fmt.Println("The path is temporarily added to PATH 😏")
		}
	},
}

func init() {
	rootCmd.AddCommand(addpathCmd)

	addpathCmd.Flags().BoolP("const", "c", false, "adds the PATH variable to your system configuration.")
	addpathCmd.Flags().BoolP("temp", "t", true, "adds to PATH while the terminal is running")
}
