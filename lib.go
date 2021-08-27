package asciinema

import (
	"os"

	"github.com/securisec/asciinema/asciicast"
	"github.com/securisec/asciinema/commands"
	"github.com/securisec/asciinema/util"
)

// Options options to pass to various commands.
// These are common flags passed to the asciinema cli.
type Options struct {
	Title   string
	MaxWait float64
	Yes     bool
	Quite   bool
}

// New creates a new Options instance.
func New(opts ...Options) *Options {
	var o Options
	if len(opts) == 0 {
		return &Options{
			Title:   "",
			MaxWait: 1.0,
			Yes:     false,
			Quite:   false,
		}
	}

	options := opts[0]

	o = options
	return &o
}

// Play plays the given asciicast. Use asciicast.Asciicast to unmarshal
// read from the asciicast file.
func (o *Options) Play(cast *asciicast.Asciicast) error {
	cmd := commands.NewPlayCommand()
	return cmd.Execute(cast, o.MaxWait)
}

// // Auth generates a local token and provides an auth URL.
// func (o *Options) Auth() error {
// 	cmd := commands.NewAuthCommand(api)
// 	return cmd.Execute()
// }

// // Upload uploads the given file to asciinema.
// func (o *Options) Upload(filename string) error {
// 	cmd := commands.NewUploadCommand(api)
// 	return cmd.Execute(filename)
// }

// Rec records the terminal and returns the asciicast and error.
func (o *Options) Rec() (asciicast.Asciicast, error) {
	command := util.FirstNonBlank(os.Getenv("SHELL"), cfg.RecordCommand())
	title := o.Title
	assumeYes := o.Yes

	if o.Quite {
		util.BeQuiet()
		assumeYes = true
	}

	maxWait := o.MaxWait
	cmd := commands.NewRecordCommand(api, env)
	return cmd.Execute(command, title, assumeYes, maxWait)
}
