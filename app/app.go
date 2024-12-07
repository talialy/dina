package app

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func Init() {
	var update cli.Command
	var install cli.Command

	cmd := &cli.Command{
		Commands: []*cli.Command{
			Update(&update),
			Install(&install),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
