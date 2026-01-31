//go:build windows
// +build windows

package workpath

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func AddUserPathInPATH(newPath string) error {
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

	for _, p := range strings.Split(PATH, ";") {
		fmt.Println("---------------------------------------------------")
		fmt.Println(p)
	}

	return nil
}

func RemoveUserPathInPath(path string) error {
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
