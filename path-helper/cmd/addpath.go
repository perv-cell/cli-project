/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
)

// addpathCmd represents the addpath command
var addpathCmd = &cobra.Command{
	Use:   "addpath",
	Short: "add our path in common PATH",
	Long: `
addpath adds your file path to the configuration
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
				err := addUserPathInPATH(path)
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

func removeUserPathInPATH(path string) error {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Environment`,
		registry.ALL_ACCESS,
	)
	if err != nil {
		return fmt.Errorf("Error registry.OpenKey: %w", err)
	}

	defer key.Close()

	PATH, _, err := key.GetStringValue("PATH")
	if err != nil {
		return fmt.Errorf("Couldn't extract data from Path: %w", err)
	}

	paths := strings.Split(PATH, ";")
	lenPaths := len(paths)
	newPaths := make([]string, 0, lenPaths)
	for _, p := range paths {
		if p != path {
			newPaths = append(newPaths, p)
			lenPaths--
		}
	}

	err = key.SetStringValue("PATH", strings.Join(newPaths, ";"))
	if err != nil {
		if lenPaths == 0 {
			return fmt.Errorf("your path is not in the PATH environment")
		}
		return err
	}

	return nil
}

// для удаления
func addUserPathInPATH(newPath string) error {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Environment`,
		registry.ALL_ACCESS,
	)
	if err != nil {
		return err
	}
	defer key.Close()

	PATH, _, err := key.GetStringValue("PATH")
	if err != nil {
		return err
	}

	paths := strings.Split(PATH, ";")
	for _, p := range paths {
		if p == newPath {
			return fmt.Errorf("the path has already been added")
		}
	}
	paths = append(paths, newPath)
	newPATH := strings.Join(paths, ";")

	return key.SetStringValue("PATH", newPATH)
}

func LookPATHenvirenment() error {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Envirenment`,
		registry.ALL_ACCESS,
	)
	if err != nil {
		return err
	}

	defer key.Close()

	PATH, _, err := key.GetStringValue("PATH")
	if err != nil {
		return err
	}

	fmt.Printf("PATH = %s", PATH)

	return nil
}

func init() {
	rootCmd.AddCommand(addpathCmd)

	addpathCmd.Flags().BoolP("const", "c", false, "adds the PATH variable to your system configuration.")
	addpathCmd.Flags().BoolP("temp", "t", true, "adds to PATH while the terminal is running")
}
