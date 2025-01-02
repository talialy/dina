package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

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
// this is usually config.toml unless the user specifies otherwise
func UserDinaConfig(path string) (*ConfigToml, error) {
	configFile := filepath.Join(path, "config.toml")
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

// DinaFolder returns the config folder for dina with stat. It returns os.ErrNotExist if nothing can be found
func DinaConfigDir() (os.FileInfo, error) {
	homeFolder, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path := strings.Join([]string{homeFolder, ".dina"}, "/")
	f, err := os.Stat(path)
	return f, err
}

// MakeBackup creates a backup of the folder passed. It goes inside /.dina/backups
func MakeBackup(dirPath string) error {
	var backupError error = nil
	_, err := os.Stat(dirPath)
	backupError = err

	return backupError
}

// GetBackups returns a list from the entries inside /.dina/backups
func GetBackups() ([]os.DirEntry, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	backups, err := os.ReadDir(strings.Join([]string{home, ".dina/backup"}, "/"))
	return backups, err
}
