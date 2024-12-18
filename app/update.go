package app

import (
	"bytes"
	"context"
	"dina/utils"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v3"
)

func Update(update *cli.Command) *cli.Command {
	update.Name = "update"
	update.Aliases = []string{"up"}
	update.Usage = "Updates the config.toml with the current folder structure"
	update.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "flatpak",
			Usage: "Just removes flatpaks from the output",
			Value: false,
		},
	}

	update.Action = func(ctx context.Context, c *cli.Command) error {

		var configFolders []utils.StowConfigToml
		configDir, err := os.ReadDir("config")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No local config folder found")
			return nil
		}
		var newStowConfig utils.StowConfigToml

		for _, dir := range configDir {
			if !dir.IsDir() {
				continue
			}

			newStowConfig.Name = dir.Name()
			currentDir := strings.Join([]string{"config", dir.Name()}, "/")

			dependenciesFile := strings.Join([]string{"config", dir.Name(), ".dependencies"}, "/")
			file, err := os.ReadFile(dependenciesFile)
			if errors.Is(err, os.ErrNotExist) {
				newStowConfig.Dependencies = nil
			}

			dependencies := strings.Split(string(file), "\n")
			newStowConfig.Dependencies = dependencies

			scriptsFiles := strings.Join([]string{currentDir, ".scripts"}, "/")
			dir, err := os.ReadDir(scriptsFiles)

			if errors.Is(err, os.ErrNotExist) {
				newStowConfig.Scripts = nil
			}

			for _, file := range dir {
				if file.Name() == "" {
					continue
				}
				newStowConfig.Scripts = append(newStowConfig.Dependencies, file.Name())
			}

			configFolders = append(configFolders, newStowConfig)
		}

		var flatpaks []string
		if c.Bool("flatpak") {
			fmt.Println("Going with flatpaks")
			cmd := exec.Command("flatpak", "list", "--app", "--columns=application")
			cmd.Stdout = nil
			cmd.Stderr = nil // shut that output
			out, err := cmd.Output()
			if errors.Is(err, exec.ErrNotFound) {
				fmt.Println("no flatpak was found. Running without it")
			}
			fmt.Println("Adding flatpaks...")
			flatpaks = strings.Split(string(out), "\n")
		}

		var buf = new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(utils.ConfigToml{
			Stow:     configFolders,
			Flatpaks: flatpaks,
		})

		if err != nil {
			log.Fatal(err)
		}
		currentFolder, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		configFile, err := os.Create(strings.Join([]string{currentFolder, "config.toml"}, "/"))
		if err != nil {
			log.Fatal(err)
		}
		defer configFile.Close()

		_, err = configFile.WriteString(buf.String())
		if err != nil {
			fmt.Println("there was an error while writting the file!")
			log.Fatal(err)
		}
		println("Everything's done 🍓")
		println(">> config.toml")
		return nil
	}
	return update
}
