package cmd

import (
	"errors"
	"os"
	"strings"
)

// GetPkg compares against the binaries inside the system to get the package manager, if it founds more than one, it defaults to the first one.
func GetPkg() (string, error) {
	var err error = nil
	path := "/usr/bin"
	packageManagers := map[string]string{
		"apt":    strings.Join([]string{path, "apt"}, "/"),
		"dnf":    strings.Join([]string{path, "dnf"}, "/"),
		"pacman": strings.Join([]string{path, "pacman"}, "/"),
		"zypper": strings.Join([]string{path, "zypper"}, "/"),
		"emerge": strings.Join([]string{path, "emerge"}, "/"),
	}

	var pkg string
	for name, path := range packageManagers {
		_, err := os.Stat(path)
		if err != nil {
			continue
		}
		pkg = name
		break
	}
	if pkg == "" {
		err = errors.New("No package manager found")
	}
	return pkg, err
}

// GetPkgOS uses the OS name to assign a package manager. If the OS isn't added yet, it will return an error. Stupid but maybe
func GetPkgOS() error {
	return nil
}
