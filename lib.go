package asciinema

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

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
	initAsciinema()
	cmd := commands.NewPlayCommand()
	return cmd.Execute(cast, o.MaxWait)
}

// Rec records the terminal and returns the asciicast and error.
func (o *Options) Rec() (asciicast.Asciicast, error) {
	initAsciinema()
	command := util.FirstNonBlank(os.Getenv("SHELL"), cfg.RecordCommand())
	title := o.Title
	assumeYes := o.Yes

	if o.Quite {
		util.BeQuiet()
		assumeYes = true
	}

	maxWait := o.MaxWait
	cmd := commands.NewRecordCommand(env)
	return cmd.Execute(command, title, assumeYes, maxWait)
}

func initAsciinema() {
	env = environment()

	if !util.IsUtf8Locale(env) {
		fmt.Println("asciinema needs a UTF-8 native locale to run. Check the output of `locale` command.")
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		showCursorBack()
		os.Exit(1)
	}()
	defer showCursorBack()

	cfg, err = util.GetConfig(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const Version = "1.2.0"

func environment() map[string]string {
	env := map[string]string{}

	for _, keyval := range os.Environ() {
		pair := strings.SplitN(keyval, "=", 2)
		env[pair[0]] = pair[1]
	}

	return env
}

func showCursorBack() {
	fmt.Fprintf(os.Stdout, "\x1b[?25h")
}

var (
	env map[string]string
	cfg *util.Config
	err error
)
