package commands

import (
	"fmt"

	"github.com/securisec/asciinema/api"
	"github.com/securisec/asciinema/util"
)

type UploadCommand struct {
	API api.API
}

func NewUploadCommand(api api.API) *UploadCommand {
	return &UploadCommand{
		API: api,
	}
}

func (c *UploadCommand) Execute(filename string) error {
	var url, warn string
	var err error

	util.WithSpinner(0, func() {
		url, warn, err = c.API.UploadAsciicast(filename)
	})

	if warn != "" {
		util.Warningf(warn)
	}

	if err != nil {
		return err
	}

	fmt.Println(url)

	return nil
}
