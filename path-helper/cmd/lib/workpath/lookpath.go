package workpath

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

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
