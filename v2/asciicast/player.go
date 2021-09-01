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

// TODO vv sleep time as lib is not right. correct in asciinema-player
func (r *AsciicastPlayer) Play(asciicast *Asciicast, maxWait float64) error {
	for _, frame := range asciicast.Stdout {
		delay := frame.Time
		if maxWait > 0 && delay > maxWait {
			delay = maxWait
		}
		time.Sleep(time.Duration(float64(time.Second) * delay))
		r.Terminal.Write(frame.EventData)
	}

	return nil
}
