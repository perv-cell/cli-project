//go:build linux
// +build linux

package workpath

import (
	"fmt"
	"os"
	"strings"
)

func AddUserPathInPATH(newPath string) error {
	PATH := os.Getenv("PATH")
	err := os.Setenv("PATH", PATH+";"+newPath)
	if err != nil {
		return err
	}
	return nil
}
func LookPATHenvirenment() error {
	PATH := os.Getenv("PATH")
	paths := strings.Split(PATH, ";")
	for _, p := range paths {
		fmt.Println("-----------------------------------------")
		fmt.Println(p)
	}
	if PATH == "" {
		return fmt.Errorf("PATH is empty")
	}
	return nil
}
func RemoveUserPathInPATH(path string) error {
	PATH := os.Getenv("PATH")
	paths := strings.Split(PATH, ";")
	lenPaths := len(paths)
	pathsReady := make([]string, 0, lenPaths)
	for _, p := range paths {
		if p != path {
			pathsReady = append(pathsReady, p)
			lenPaths--
		}
	}
	if lenPaths > 0 {
		return fmt.Errorf("path %s already remove", path)
	}
	return nil
}
