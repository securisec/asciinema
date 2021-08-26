package asciinema

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"

	asciinemaapi "github.com/securisec/asciinema/api"
	"github.com/securisec/asciinema/util"
)

const Version = "1.2.0"

var usage = `Record and share your terminal sessions, the right way.

Usage:
  asciinema rec [-c <command>] [-t <title>] [-w <sec>] [-y] [-q] [<filename>]
  asciinema play [-w <sec>] <filename>
  asciinema upload <filename>
  asciinema auth
  asciinema -h | --help
  asciinema --version

Commands:
  rec             Record terminal session
  play            Replay terminal session
  upload          Upload locally saved terminal session to asciinema.org
  auth            Assign local API token to asciinema.org account

Options:
  -c, --command=<command>  Specify command to record, defaults to $SHELL
  -t, --title=<title>      Specify title of the asciicast
  -w, --max-wait=<sec>     Reduce recorded terminal inactivity to max <sec> seconds (can be fractional)
  -y, --yes                Answer "yes" to all prompts (e.g. upload confirmation)
  -q, --quiet              Be quiet, suppress all notices/warnings (implies -y)
  -h, --help               Show this message
  --version                Show version`

func cmdName(args map[string]interface{}) string {
	for _, cmd := range []string{"rec", "play", "upload", "auth"} {
		if args[cmd].(bool) {
			return cmd
		}
	}

	return ""
}

func stringArg(args map[string]interface{}, name string) string {
	val := args[name]

	if val != nil {
		return val.(string)
	} else {
		return ""
	}
}

func boolArg(args map[string]interface{}, name string) bool {
	return args[name].(bool)
}

func floatArg(args map[string]interface{}, name string, defaultValue float64) float64 {
	val := args[name]

	if val != nil {
		number, err := strconv.ParseFloat(val.(string), 64)

		if err == nil {
			return float64(number)
		}
	}

	return defaultValue
}

func formatVersion() string {
	return fmt.Sprintf("asciinema %v", Version)
}

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
	api *asciinemaapi.AsciinemaAPI
	err error
)

func init() {
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
	api = asciinemaapi.New(cfg.ApiUrl(), env["USER"], cfg.ApiToken(), Version)
}

// CmdLineTool is a command line tool that can be used to record terminal sessions.
// This function was originall main function in asciinema repo.
// Some of the internal logic was moved to the init() function.
// func CmdLineTool() {
// 	args, _ := docopt.Parse(usage, nil, true, formatVersion(), false)

// 	switch cmdName(args) {
// 	case "rec":
// 		command := util.FirstNonBlank(stringArg(args, "--command"), cfg.RecordCommand())
// 		title := stringArg(args, "--title")
// 		assumeYes := cfg.RecordYes() || boolArg(args, "--yes")

// 		if boolArg(args, "--quiet") {
// 			util.BeQuiet()
// 			assumeYes = true
// 		}

// 		maxWait := floatArg(args, "--max-wait", cfg.RecordMaxWait())
// 		filename := stringArg(args, "<filename>")
// 		cmd := commands.NewRecordCommand(api, env)
// 		err = cmd.Execute(command, title, assumeYes, maxWait, filename)

// 	case "play":
// 		maxWait := floatArg(args, "--max-wait", cfg.PlayMaxWait())
// 		filename := stringArg(args, "<filename>")
// 		cmd := commands.NewPlayCommand()
// 		err = cmd.Execute(filename, maxWait)

// 	case "upload":
// 		filename := stringArg(args, "<filename>")
// 		cmd := commands.NewUploadCommand(api)
// 		err = cmd.Execute(filename)

// 	case "auth":
// 		cmd := commands.NewAuthCommand(api)
// 		err = cmd.Execute()
// 	}

// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(1)
// 	}
// }
