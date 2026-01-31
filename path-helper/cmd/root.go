/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:       "pelp",
	Short:     "The commands are:",
	Long:      `"pelp is a tool for work ENV, PATH"`,
	ValidArgs: []string{"addpath", "rmpath", "lookpath"},
	Args:      cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Возвращаем список ресурсов, которые можно листать.
		// toComplete - это часть строки, которую пользователь уже начал вводить.
		resources := []string{"pods", "services", "deployments", "configmaps"}
		// Фильтруем по тому, что пользователь уже ввел (toComplete)
		var completions []string
		for _, r := range resources {
			if strings.HasPrefix(r, toComplete) {
				completions = append(completions, r)
			}
		}
		// Можно вернуть с директивой, например, чтобы указать, что это не имена файлов
		return completions, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("MY error")
	},
}

func NewExecute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SilenceUsage = false  // если true то usage не будет выводиться
	rootCmd.SilenceErrors = false // если true то ошибки не выводяться программа просто завершается
	rootCmd.AddGroup(&cobra.Group{
		ID:    "path",
		Title: "Path commands",
	})
}
