package main

import (
	"context"
	"fmt"
	"gitpusher/pkg/commander"
	"gitpusher/pkg/config"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// ctx := context.Background()
	// c := &config.RunConfig{}
	// if err := run(ctx, c, os.Stdout); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "savefile",
				Aliases: []string{"c"},
				Usage:   "commit a file and push to remote",
				Action: func(c *cli.Context) error {
					var commitCmd = []string{"git", "commit", "-a"}

					commitCmd = append(commitCmd, "-m")
					commitCmd = append(commitCmd, "debug")

					// testCmd := []string{"tr", "a-z", "A-Z"}

					cmd := commander.Command{Parts: commitCmd}
					cmd.RunCommand()

					pushCmd := []string{"git", "push", "origin", "master"}

					pshCmd := commander.Command{Parts: pushCmd}
					pshCmd.RunCommand()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func run(ctx context.Context, c *config.RunConfig, out io.Writer) error {
	c.Init(os.Args)
	log.SetOutput(out)

	var commitCmd = []string{"git", "commit", "-a"}

	if c.Sign {
		commitCmd = append(commitCmd, "-S")
	}

	commitCmd = append(commitCmd, "-m")
	commitCmd = append(commitCmd, c.CommitMessage)

	// testCmd := []string{"tr", "a-z", "A-Z"}

	cmd := commander.Command{Parts: commitCmd}
	cmd.RunCommand()

	pushCmd := []string{"git", "push", "origin", "master"}

	pshCmd := commander.Command{Parts: pushCmd}
	pshCmd.RunCommand()

	return nil
}
