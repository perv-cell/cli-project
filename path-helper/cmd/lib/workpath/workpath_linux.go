//go:build linux
// +build linux

package workpath

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func AddUserPathInPATH(newPath string) error {

	pathConfig, err := findFileConfigInUserDirectory()
	if err != nil {
		return err
	}
	err = _addPathInConfigFile(pathConfig, newPath)
	if err != nil {
		return err
	}
	return nil

}

func LookPATHenvirenment() error {
	PATH := os.Getenv("PATH")
	paths := strings.Split(PATH, ":")
	for _, p := range paths {
		fmt.Println("-----------------------------------------")
		fmt.Println(p)
	}
	if PATH == "" {
		return fmt.Errorf("PATH is empty")
	}
	return nil
}

func RemoveUserPathInPath(path string) error {

	pathConfig, err := findFileConfigInUserDirectory()
	if err != nil {
		return err
	}
	err = _removePathFromConfig(pathConfig, path)
	if err != nil {
		return err
	}
	return nil
}

func _removePathFromConfig(pathToConfig, path string) error {
	content, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	patterns := []string{
		fmt.Sprintf("export PATH=$PATH:%s\n", path),
		fmt.Sprintf("export PATH=$PATH:%s\r\n", path),
		fmt.Sprintf("\nexport PATH=$PATH:%s\n", path),
		fmt.Sprintf("PATH=$PATH:%s\n", path),
		fmt.Sprintf("PATH=$PATH:%s;", path),
	}

	var found bool
	newContent := content

	for _, pattern := range patterns {
		idx := bytes.Index(newContent, []byte(pattern))
		if idx != -1 {
			newContent = append(newContent[:idx], newContent[idx+len(pattern):]...)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("path %q is not found in PATH", path)
	}

	err = os.WriteFile(pathToConfig, newContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	fmt.Printf("Путь %s успешно удалён в %s\n", path, pathToConfig)
	fmt.Println("Чтобы изменения вступили в силу, выполните команду:")
	fmt.Printf("source %s\n", pathToConfig)

	return nil
}

func _addPathInConfigFile(pathToConfig, path string) error {

	content, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("Error: failed read config: %w", err)
	}

	pathEntry := fmt.Sprintf("\nexport PATH=$PATH:%s\n", path)
	if strings.Contains(string(content), pathEntry) {
		fmt.Printf("Путь %s уже есть в файле %s\n", path, pathToConfig)
		return nil
	}
	f, err := os.OpenFile(pathToConfig, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(pathEntry); err != nil {
		return err
	}
	fmt.Printf("Путь %s успешно добавлен в %s\n", path, pathToConfig)
	fmt.Println("Чтобы изменения вступили в силу, выполните команду:")
	fmt.Printf("source %s\n", pathToConfig)

	return nil
}

func findFileConfigInUserDirectory() (string, error) {
	listMaybeConfigFiles := []string{".bashrc", ".profile", ".zshrc"}

	bld := strings.Builder{}
	pathToUser, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error: %w. Path to user not founded.", err)
	}

	var pathConfig string
	for _, fConfig := range listMaybeConfigFiles {
		bld.WriteString(pathToUser)
		bld.WriteString("/")
		bld.WriteString(fConfig)
		pathConfig = bld.String()
		if _, err := os.Stat(pathConfig); !os.IsNotExist(err) {
			break
		}
		bld.Reset()
		pathConfig = ""
	}
	if pathConfig == "" {
		return "", fmt.Errorf(`
		Error: %w. Not founded file config. File config for example: .bashrc, .profile, .zshrc in User directory`, err)
	}

	return pathConfig, nil
}
