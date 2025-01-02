package app

import (
	"os"
	"strings"
)

type Filetree struct {
	config   os.DirEntry
	external []os.DirEntry
	scripts  os.DirEntry
	apps     os.DirEntry
}

func CreateTree(dirs []os.DirEntry) (Filetree, error) {
	tree := Filetree{
		config:   nil,
		external: nil,
		scripts:  nil,
		apps:     nil,
	}

	for _, folder := range dirs {
		if folder.IsDir() {
			switch folder.Name() {
			case ".config":
				tree.config = folder
			default:
				if strings.Contains(folder.Name(), ".") {
					tree.external = append(tree.external, folder)
				}
			}
		}
	}
	return tree, nil
}
