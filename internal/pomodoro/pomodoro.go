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
	for i := int(minutes) * 60; i >= 0; i-- {
		select {
		case <-p.ctx.Done():
			return
		default:
			fmt.Print("\033[s\033[K")
			fmt.Printf("\033[48;5;220m") // set foreground
			fmt.Printf("\033[38;5;16m")  // set background
			fmt.Printf(" %s ", message)
			fmt.Printf("\033[0m")        // reset colors
			fmt.Printf("\033[48;5;16m")  // set foreground
			fmt.Printf("\033[38;5;220m") // set background
			fmt.Printf(" %d:%02d\033[u ", i/60, i%60)
			fmt.Printf("\033[0m ")

			time.Sleep(time.Second)
		}
	}
}
