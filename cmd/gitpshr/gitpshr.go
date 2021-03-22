package main

import (
	"fmt"
	"gitpusher/pkg/promter"

	"github.com/namsral/flag"
)

func main() {
	p := promter.NewService()
	var prompt promter.Prompt
	prompt.Text = "Hello World\n"
	p.Prompt(prompt)
	var age int
	flag.IntVar(&age, "age", 0, "age of gopher")
	flag.Parse()
	fmt.Print("age:", age)
}
