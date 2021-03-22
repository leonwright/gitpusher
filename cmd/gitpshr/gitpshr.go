package main

import (
	"gitpusher/pkg/promter"
)

func main() {
	p := promter.NewService()
	var prompt promter.Prompt
	prompt.Text = "Hello World\n"
	p.Prompt(prompt)
}
