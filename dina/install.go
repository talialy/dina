package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/talialy/dina/cmd/app"
	"github.com/talialy/dina/cmd/backups"
	"github.com/talialy/dina/utils/config"
	"github.com/urfave/cli/v3"
)

func Install(install *cli.Command) *cli.Command {
	install.Name = "install"
	install.Aliases = []string{"ins", "i"}
	install.Usage = "Uses config.toml to setup the system"
	install.Description = "Install your configuration files and packages in a single command. If multiple flags are passed, only the safest one will be used (default: omit)"

	install.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "omit",
			Aliases: []string{"o"},
			Usage:   "If a folder conflicts, it's omitted",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Deletes and then links every conflicting folder",
		},
		&cli.BoolFlag{
			Name:    "backup",
			Aliases: []string{"b"},
			Usage:   "Creates a backup of every folder conflicting",
		},
		&cli.StringFlag{
			Name:  "backup-name",
			Usage: "Instead of using a UTC stamp, it uses this for the backup folder",
		},
		&cli.StringFlag{
			Name:  "config",
			Usage: "change the default config file path for this one",
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

		configFilePath := "./"
		if c.String("config") != "" {
			absolute, err := filepath.Abs(c.String("config"))
			if err != nil {
				log.Fatal(err)
			}
			configFilePath = absolute
		}

		configFile, err := config.ReadConfFile(configFilePath)
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("no config file found!")
		}

		configPath := config.UserConfigDir()
		opts := app.OptHandlerFlags{}
		switch {
		case c.Bool("omit"):
			opts.Omit = true
		case c.Bool("backup"):
			opts.Backup = true
			timeNow := time.Now().Format("2006-01-02_15-04-05")
			opts.BackupStamp = strings.Join([]string{"backup", timeNow}, "_")
			if _, err := backups.GetBckp(opts.BackupStamp); !errors.Is(err, os.ErrNotExist) {
				log.Fatal(err)
			}
		case c.Bool("force"):
			opts.Force = true
		}

	installer:
		for _, folder := range configFile.Stow {
			targetPath := filepath.Join("config", folder.Name)
			workingFolder, err := filepath.Abs(targetPath)
			if err != nil {
				log.Fatal(err)
			}
			_, errStatDir := os.Stat(workingFolder)
			if errors.Is(errStatDir, os.ErrNotExist) {
				log.Fatalf("The folder %s inside .dina.toml doesn't exist, check again\n", workingFolder)
			}

			targetFolder := filepath.Join(configPath, folder.Name)
			_, err = os.Stat(targetFolder)
			folderExist := !errors.Is(err, os.ErrNotExist)
			if folderExist {
				control := app.FilesHandler(opts, folder.Name)
				switch control {
				case app.Continue:
					continue installer
				case app.Break:
					break installer
				}
			}
			err = os.Symlink(workingFolder, targetFolder)
			if err != nil {
				log.Fatal("there was an error while linking the folders ", err)
			}
		}
		return nil
	}
	return install
}
