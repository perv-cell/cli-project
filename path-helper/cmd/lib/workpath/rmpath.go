package workpath

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func RemoveUserPathInPATH(path string) error {
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
