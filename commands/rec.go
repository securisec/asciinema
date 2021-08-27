package commands

import (
	"io/ioutil"

	"github.com/securisec/asciinema/api"
	"github.com/securisec/asciinema/asciicast"
)

type RecordCommand struct {
	API      api.API
	Env      map[string]string
	Recorder asciicast.Recorder
}

func NewRecordCommand(api api.API, env map[string]string) *RecordCommand {
	return &RecordCommand{
		API:      api,
		Env:      env,
		Recorder: asciicast.NewRecorder(),
	}
}

func (c *RecordCommand) Execute(command, title string, assumeYes bool, maxWait float64) (asciicast.Asciicast, error) {
	// var upload bool
	// var err error

	// if filename != "" {
	// 	upload = false
	// } else {
	// 	filename, err = tmpPath()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	upload = true
	// }

	return c.Recorder.Record(command, title, maxWait, assumeYes, c.Env)

	// if upload {
	// 	if !assumeYes {
	// 		util.Printf("Press <Enter> to upload, <Ctrl-C> to cancel.")
	// 		util.ReadLine()
	// 	}

	// 	var url, warn string
	// 	var err error

	// 	util.WithSpinner(0, func() {
	// 		url, warn, err = c.API.UploadAsciicast(filename)
	// 	})

	// 	if warn != "" {
	// 		util.Warningf(warn)
	// 	}

	// 	if err != nil {
	// 		util.Warningf("Upload failed, asciicast saved at %v", filename)
	// 		util.Warningf("Retry later by executing: asciinema upload %v", filename)
	// 		return err
	// 	}

	// 	os.Remove(filename)
	// 	fmt.Println(url)
	// }

	// return nil
}

func tmpPath() (string, error) {
	file, err := ioutil.TempFile("", "asciicast-")
	if err != nil {
		return "", err
	}
	defer file.Close()

	return file.Name(), nil
}
