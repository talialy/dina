package backups

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/talialy/dina/utils"
	"github.com/talialy/dina/utils/config"
)

// errBackupExist portable error
var ErrBackupExist error = errors.New("the folder that it tried to create was already there")

type BackupFolder struct {
	stamp string
	list  []os.DirEntry
	path  string
}

func GetPath() string {
	confFolder := config.GetPath()
	return filepath.Join(confFolder, "backups")
}

// Create s a folder inside ~/.dina/backup using stamp
// as the name of the folder. If there is a folder with the
// same stamp, it will pass os.ErrExist error
func Create(stamp string) error {
	bckpPath := GetPath()

	path := filepath.Join(bckpPath, stamp)
	err := os.MkdirAll(path, 0755)
	if errors.Is(err, os.ErrExist) {
		return ErrBackupExist
	}
	return nil
}

// ReadDirBckp returns a BackupFolder structure with
// the backup stamp, if the folder doesnt exists, returns
// os.ErrNotExist error
func ReadDirBckp(stamp string) (BackupFolder, error) {
	info := BackupFolder{
		stamp: stamp,
	}
	var errReadDir error = nil

	bckpPath := GetPath()

	info.path = filepath.Join(bckpPath, stamp)
	list, err := os.ReadDir(info.path)
	if errors.Is(err, os.ErrNotExist) {
		info.list = nil
		errReadDir = os.ErrNotExist
	}
	info.list = list
	return info, errReadDir
}

var ErrTargetNotExist error = errors.New("the folder or file that you tried to backup didn't exist")
var ErrBckpNotExist error = errors.New("the backup you tried to access doesn't exist!")

// MkBckp takes a folder or a file and stores it inside ~/.dina
// taking stamp as a identifier on where to store it.
//
// This should normally be a timestamp using UTC unless directed by the user
func MkBckp(dir string, stamp string) error {
	var backupError error = nil

	_, err := os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		return ErrTargetNotExist
	}

	path, err := GetBckp(stamp)
	if errors.Is(err, os.ErrNotExist) {
		return ErrBckpNotExist
	}
	dirName := filepath.Base(dir)
	newDirLocation := filepath.Join(path, dirName)

	_, err = utils.CopyDir(dir, newDirLocation)
	if err != nil {
		log.Fatal(err)
	}
	return backupError
}

// Get returns the folder requested from ~/.dina/backups
func GetBckp(stamp string) (string, error) {
	var errGet error = nil

	bckpPath := GetPath()
	path := filepath.Join(bckpPath, stamp)
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		errGet = os.ErrNotExist
	} else if err != nil {
		log.Fatal(err)
	}
	return path, errGet
}
