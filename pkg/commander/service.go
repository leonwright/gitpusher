package commander

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (c *Command) RunCommand() {
	cmd := exec.Command(c.Parts[0], c.Parts[1:]...)
	// fmt.Println(cmd.String())

	//cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	os.Stdout.Write(out.Bytes())

	err := cmd.Run()

	if err != nil {
		fmt.Println(stdErr.String())
		log.Fatal(err.Error())
	}

	os.Stdout.Write(out.Bytes())
}
