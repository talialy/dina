package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func Init() { 
    cmd := &cli.Command{
	Commands: []*cli.Command{
	    {
		Name: "update",
		Aliases: []string{"up"},
		Action: func(ctx context.Context, c *cli.Command) error {
		    dir, err := os.Getwd();
		    if err != nil {
			return err
		    }
		    list, _ := os.ReadDir(dir) 
		    for _, folder := range list {
		    }
		    return nil
		},
	    },
	},
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }

}
