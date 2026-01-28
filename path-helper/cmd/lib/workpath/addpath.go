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
