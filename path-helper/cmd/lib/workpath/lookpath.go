package workpath

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

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
