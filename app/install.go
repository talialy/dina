package app

import (
	"context"
	"dina/utils"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v3"
)

func Install(install *cli.Command) *cli.Command {
	install.Name = "install"
	install.Aliases = []string{"ins"}
	install.Usage = "Uses config.toml to setup the system"
	install.Description = "Using the config inside the current directory, it goes trough all the options to install the files and applications to the system"

	install.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "overwrites all folders ignoring the contents (without backup)",
		},
		&cli.BoolFlag{
			Name:    "omit",
			Aliases: []string{"o"},
			Usage:   "install what it can, omits if the folder is found",
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

		read, err := os.ReadFile("config.toml")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("config.toml not found! :(")
		}

		var config utils.ConfigToml
		_, err = toml.Decode(string(read), &config)

		configPath, err := os.UserConfigDir()
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Stat(configPath)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Why no config folder?")
			return nil
		}

		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		for _, folder := range config.Stow {
			TargetFolder := strings.Join([]string{configPath, folder.Name, ""}, "/")
			currentFolder := strings.Join([]string{currentDirectory, "config", folder.Name, ""}, "/")
			_, err := os.Stat(TargetFolder)

			if !errors.Is(err, os.ErrNotExist) {
				switch {
				case c.Bool("omit"):
					println(folder.Name, "exists, omitting")
				case c.Bool("force"):
					println(folder.Name, "exists, forcing contents")
				default:
					println(folder.Name, "folder already exists!")
					println("if you really want to continue")
					println("run with the --force flag or --omit")
					return nil
				}
			}

			cmd := exec.Command("ln", "-s", currentFolder, configPath)
			fmt.Println("ln", "-s", currentFolder, configPath)
			cmd.Stdout = nil
			cmd.Stderr = nil // shut that output, yeah, rehusing code
			_, err = cmd.Output()
			if err != nil {
				log.Fatal(err)
			}

		}
		return nil
	}
	return install
}
