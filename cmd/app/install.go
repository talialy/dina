package app

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/talialy/dina/utils/config"
)

// OptHandlerFlags gives the mode that the files are going to be installing.
// it depends on the flag --type withing dina install
type OptHandlerFlags struct {
	Omit   bool
	Backup bool
	Force  bool
}

// controlFlow is used to indicate a higher for loop on what to do next
type controlFlow int

const (
	void controlFlow = iota
	Continue
	Break
)

// FilesHandler handles each case within --type options.
// If no option to handle the files is passed, it will return the state of the folder
// between os.ErrExist and nil
//
// opt.Omit
// ignores the error
// opt.Force
// removes or unlinks the folder with os.Remove()
// opt.Backup
// creates a timestap backup inside ~/.dina and moves the folder inside it
func FilesHandler(opt OptHandlerFlags, folder string) controlFlow {
	configPath := config.UserConfigDir()

	targetFolder := strings.Join([]string{configPath, folder}, "/")
	switch {
	case opt.Omit:
		fmt.Printf("%s was found, skipping\n", folder)
		return Continue
	case opt.Backup:
		fmt.Printf("making backup of %s and installing", folder)
		err := config.MakeBackup(targetFolder)
		if err != nil {
			log.Fatal("there was an error while doing the backup", err)
		}
		err = os.Remove(targetFolder)
		if err != nil {
			log.Fatal("there was an error while removing the folder", err)
		}
		return void
	case opt.Force:
		// scary!
		fmt.Printf("%s was found, forcesfully installing\n", folder)
		info, err := os.Lstat(targetFolder)
		if err != nil {
			log.Fatal("there was an error reading the folder", err)
		}
		switch {
		case info.Mode()&os.ModeSymlink != 0:
			err := syscall.Unlink(targetFolder)
			if err != nil {
				log.Fatal(err)
			}
		case info.IsDir():
			err = os.Remove(targetFolder)
			if err != nil {
				log.Fatal("there was an error while removing the folder ", err)
			}
		default:
			log.Fatal("I don't know how to handle this file . .)")
		}
	default:
		fmt.Println("One or more folders are already present. Use --type=omit to ignore or use another flag")
		return Break
	}
	return void
}
