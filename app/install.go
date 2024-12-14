package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"momo/utils"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v3"
)

func Install(install *cli.Command) *cli.Command {
	install.Name = "install"
	install.Aliases = []string{"ins"}
	install.Usage = "Uses config.toml to setup the system"
	install.Description = "Uses the folder structure to generate an array of packages to install using your package manager. \n If no command passed, it will use both flatpaks and snaps to get the packages"
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

		read, err := os.ReadFile("config.toml")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("config.toml not found! :(")
		} else if err != nil {
			log.Fatal(err)
		}

		var config utils.ConfigToml
		_, err = toml.Decode(string(read), &config)
		configPath, err := os.UserConfigDir()
		if err != nil {
			log.Fatal(err)
		}

		for _, folder := range config.Stow {
			selectedFolder := strings.Join([]string{configPath, folder}, "/")
			_, err := os.ReadDir(selectedFolder)

			if err != nil {
				return nil
			}

			if !c.Bool("omit") {
				println(folder, "folder already exists!")
				println("if you really want to continue")
				println("run with the --force flag or --omit")
				return nil
			} else {
				println(folder, "exists, omitting")
			}

		}
		return nil
	}
	return install
}
