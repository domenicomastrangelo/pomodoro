package pomodoro

import (
	"context"
	"fmt"
	"time"
)

type Pomodoro struct {
	PomodoroMinutes   uint
	ShortBreakMinutes uint
	LongBreakMinutes  uint
	ctx               context.Context
}

func New(ctx context.Context) *Pomodoro {
	return &Pomodoro{
		PomodoroMinutes:   25,
		ShortBreakMinutes: 5,
		LongBreakMinutes:  15,
		ctx:               ctx,
	}
}

func (p *Pomodoro) Start() {
	p.notifyCountdown(p.PomodoroMinutes, "POMODORO")
}

func (p *Pomodoro) ShortBreak() {
	p.notifyCountdown(p.ShortBreakMinutes, "SHORT BREAK")
}

func (p *Pomodoro) LongBreak() {
	p.notifyCountdown(p.LongBreakMinutes, "LONG BREAK")
}

func (p *Pomodoro) notifyCountdown(minutes uint, message string) {
	seconds := minutes * 60

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-ticker.C:
			if seconds == 0 {
				return
			}
			fmt.Printf("\033[s\033[K")
			fmt.Printf("\033[48;5;220m") // set foreground
			fmt.Printf("\033[38;5;16m")  // set background
			fmt.Printf(" %s ", message)
			fmt.Printf("\033[0m")        // reset colors
			fmt.Printf("\033[48;5;16m")  // set foreground
			fmt.Printf("\033[38;5;220m") // set background
			fmt.Printf(" %d:%02d ", seconds/60, seconds%60)
			fmt.Printf("\033[0m\033[u")

			seconds--
		default:
		}
	}
}
