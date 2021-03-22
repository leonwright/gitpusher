package promter

import (
	"fmt"
)

type Service interface {
	Prompt(p Prompt)
	PromptForBool(p Prompt) bool
	PromptForString(p Prompt) string
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Prompt(p Prompt) {
	fmt.Print(p.Text)
}

func (s *service) PromptForBool(p Prompt) (response bool) {
	s.Prompt(p)
	return
}

func (s *service) PromptForString(p Prompt) (response string) {
	s.Prompt(p)
	return
}
