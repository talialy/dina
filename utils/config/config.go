package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type StowConfigToml struct {
	Name         string
	Dependencies []string
	Scripts      []string
}
type ConfigToml struct {
	Stow     []StowConfigToml
	Flatpaks []string
}

// UserConfigDir is a wrapper for os.UserConfigDir with a custom message for the user.
func UserConfigDir() string {
	path, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {

		log.Fatal("Why no config folder?")
	}
	return path
}

// it returns the installer config for dina.
// this is usually config.toml unless the user specifies otherwise.
// If nothing can be found it returns os.ErrNotExist
func ReadConfFile(path string) (*ConfigToml, error) {
	configFile := filepath.Join(path, ".dina.toml")
	read, err := os.ReadFile(configFile)
	if errors.Is(err, os.ErrNotExist) {
		return nil, os.ErrNotExist
	}

	var config ConfigToml
	_, err = toml.Decode(string(read), &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config, nil
}

// DinaFolder returns the path to the config folder for dina.
// If nothing is found, it will create the folder before
// returning the path
func GetPath() string {
	homeFolder, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(homeFolder, ".dina")
	err = os.Mkdir(path, 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}
	return path
}
