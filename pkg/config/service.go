package config

import "flag"

type RunConfig struct {
	CommitMessage string
	Push          bool
	Sign          bool
	TagCommit     bool
	TagLabel      string
}

func (c *RunConfig) Init(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		commitMessage = flags.String("m", "", "commit message")
		push          = flags.Bool("p", true, "push to origin")
		sign          = flags.Bool("s", false, "sign the commit")
		tag           = flags.Bool("t", true, "tag the commit")
		tagLabel      = flags.String("tl", "", "label the tag (only required if -t true)")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	c.CommitMessage = *commitMessage
	c.Push = *push
	c.Sign = *sign
	c.TagCommit = *tag
	c.TagLabel = *tagLabel

	return nil
}
