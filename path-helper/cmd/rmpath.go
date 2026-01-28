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

// rmpathCmd represents the rmpath command
var rmpathCmd = &cobra.Command{
	Use:   "rmpath",
	Short: "remove our path in common PATH",
	Long: `
removes the path from the configuration file
or temporary storage, during a session
of working with the system environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("To run this command, specify the path you want to add to PATH as the only argument.")
			return
		}
		path := args[0]

		_const, _ := cmd.Flags().GetBool("const")
		if _const {
			os := runtime.GOOS
			switch os {

			case "windows":
				err := workpath.RemoveUserPathInPATH(path)
				if err != nil {
					fmt.Println(err)
				}

			case "linux":
				fmt.Println("Запущено на Linux")

			case "darwin":
				fmt.Println("Запущено на macOS (Darwin)")

			default:
				fmt.Printf("Другая ОС: %s\n", os)
			}

			fmt.Println("the path was successfully deleted 🙃")
		} else {

			fmt.Println("the path was successfully deleted 🙃")
		}
	},
}

func init() {
	rootCmd.AddCommand(rmpathCmd)

	rmpathCmd.Flags().BoolP("temp", "t", false, "Help message for toggle")
	rmpathCmd.Flags().BoolP("const", "c", false, "Help message for toggle")
}
