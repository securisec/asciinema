package asciicast

import (
	"time"

	"github.com/securisec/asciinema/v2/terminal"
)

type Player interface {
	Play(*Asciicast, float64) error
}

type AsciicastPlayer struct {
	Terminal terminal.Terminal
}

func NewPlayer() Player {
	return &AsciicastPlayer{Terminal: terminal.NewTerminal()}
}

func (r *AsciicastPlayer) Play(asciicast *Asciicast, maxWait float64) error {
	lenFrames := len(asciicast.Stdout)
	for i, frame := range asciicast.Stdout {
		delay := frame.Time
		if i < lenFrames-1 {
			delay = asciicast.Stdout[i+1].Time - delay
		} else {
			delay = asciicast.Stdout[i].Time - delay
		}
		if maxWait > 0 && delay > maxWait {
			delay = maxWait
		}
		time.Sleep(time.Duration(float64(time.Second) * delay))
		r.Terminal.Write(frame.EventData)
	}

	return nil
}
