package main

import (
	"context"
	"fmt"
	"os"

	ps "code"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			result, err := ps.GetPathSize(cmd.Args().Get(0), false, false, false)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
