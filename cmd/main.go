package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/domenicomastrangelo/pomodoro/internal/pomodoro"
)

func main() {
	pomodoroMinutes := flag.Uint("m", 25, "minutes of pomodoro")
	shortBreakMinutes := flag.Uint("s", 5, "minutes of short break")
	longBreakMinutes := flag.Uint("l", 15, "minutes of long break")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	// Cancel context on Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()

	p := pomodoro.New(ctx)
	p.PomodoroMinutes = *pomodoroMinutes
	p.ShortBreakMinutes = *shortBreakMinutes
	p.LongBreakMinutes = *longBreakMinutes

	pomodoroCount := 0

	for {
		pomodoroCount++

		select {
		case <-ctx.Done():
			return
		default:
		}

		p.Start()

		if pomodoroCount%4 == 0 {
			p.LongBreak()
		} else {
			p.ShortBreak()
		}
	}
}
