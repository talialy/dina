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
	"time"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v3"
)

func Update(update *cli.Command) *cli.Command {
	update.Name = "update"
	update.Aliases = []string{"up"}
	update.Usage = "Updates the config.toml with the current folder structure"
	update.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "no-flatpak",
			Usage: "Just removes flatpaks from the output",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "stow-comp",
			Usage: "Creates a file that is stow readable",
			Value: false,
		},
	}
	update.Action = func(ctx context.Context, c *cli.Command) error {

		fmt.Println("getting stuff ready...")

		var configFolders []string
		configDir, err := os.ReadDir("config")
		if os.IsNotExist(err) {
			fmt.Println("No config folder :|")
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, dir := range configDir {
			if dir.IsDir() {
				configFolders = append(configFolders, dir.Name())
			}
		}

		var flatpaks []string
		if !c.Bool("no-flatpak") && !c.Bool("only-stow") {
			fmt.Println("Going with flatpaks")
			cmd := exec.Command("flatpak", "list", "--app", "--columns=application")
			cmd.Stdout = nil
			cmd.Stderr = nil // shut that output
			out, err := cmd.Output()
			if errors.Is(err, exec.ErrNotFound) {
				fmt.Println("no flatpak was found. Running without it")
				time.Sleep(300)
			} else if err != nil {
				log.Fatal(err)
			}
			flatpaks = strings.Split(string(out), "\n")

			var buf = new(bytes.Buffer)
			err = toml.NewEncoder(buf).Encode(utils.ConfigToml{
				Stow:     configFolders,
				Flatpaks: flatpaks,
			})
			if err != nil {
				log.Fatal(err)
			}

			configFile, err := os.OpenFile("config.toml", os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer configFile.Close()

			_, err = configFile.WriteString(buf.String())
			if err != nil {
				fmt.Println("there was an error while writting the file!")
				log.Fatal(err)
			}

			println("Output to config.toml")
		}
		return nil
	}
	return update
}
