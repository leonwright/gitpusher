package promter

import (
	"fmt"
)

type Service interface {
	Prompt(p Prompt)
	PromptForBool(p Prompt) bool
	PromptForString(p Prompt) string
}

func (p *Prompt) Prompt() {
	fmt.Print(p.Text)
}

func (p *Prompt) PromptForBool() (response bool) {
	p.Prompt()
	return
}

func (p *Prompt) PromptForString() (response string) {
	p.Prompt()
	return
}
