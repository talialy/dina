package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/talialy/dina/cmd/app"
	"github.com/talialy/dina/utils/config"
	"github.com/urfave/cli/v3"
)

func Install(install *cli.Command) *cli.Command {
	install.Name = "install"
	install.Aliases = []string{"ins", "i"}
	install.Usage = "Uses config.toml to setup the system"
	install.Description = "Using the config inside the current directory, it goes trough all the options to install the files and applications to the system"

	install.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   "--type=omit,backup,force.\nchanges how the installer treats folders inside ~/.config/",
		},
	}

	install.Action = func(ctx context.Context, c *cli.Command) error {
		us, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		if us.Username == "root" {
			fmt.Println("Do not run as root")
			return nil
		}

		configFile, err := config.UserDinaConfig("./")
		if err != nil {
			cli.Exit("config.toml not found! :(", 1)
		}
		configPath := config.UserConfigDir()

		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		opts := app.OptHandlerFlags{}
		switch c.String("type") {
		case "backup":
			opts.Backup = true
		case "omit":
			opts.Omit = true
		case "force":
			opts.Force = true
		}

	installer:
		for _, folder := range configFile.Stow {
			currentFolder := strings.Join([]string{currentDirectory, "config", folder.Name, ""}, "/")
			targetFolder := strings.Join([]string{configPath, folder.Name}, "/")
			_, err := os.Stat(targetFolder)

			if !errors.Is(err, os.ErrNotExist) {
				control := app.FilesHandler(opts, folder.Name)
				switch control {
				case app.Continue:
					continue installer
				case app.Break:
					break installer
				}
			}

			err = os.Symlink(currentFolder, targetFolder)
			if err != nil {
				log.Fatal("there was an error while linking the folders ", err)
			}

		}
		return nil
	}
	return install
}
