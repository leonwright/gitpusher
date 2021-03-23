package main

import (
	"context"
	"fmt"
	"gitpusher/pkg/commander"
	"gitpusher/pkg/config"
	"io"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	c := &config.RunConfig{}
	if err := run(ctx, c, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
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

	return nil
}
