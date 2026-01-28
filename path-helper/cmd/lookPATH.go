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

// lookPATHCmd represents the lookPATH command
var lookPATHCmd = &cobra.Command{
	Use:   "lookpath",
	Short: "look in our PATH",
	Long:  `allows you to view the contents of the PATH`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("To run this command, specify the path you want to add to PATH as the only argument.")
			return
		}
		os := runtime.GOOS
		switch os {
		case "windows":
			err := workpath.LookPATHenvirenment()
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
	},
}

func init() {
	rootCmd.AddCommand(lookPATHCmd)

}
