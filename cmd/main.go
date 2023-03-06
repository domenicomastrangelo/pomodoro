package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"os"
	"os/signal"

	"github.com/domenicomastrangelo/pomodoro/internal/pomodoro"
)

func main() {
	pomodoroMinutes := flag.Uint("m", 25, "minutes of pomodoro")
	shortBreakMinutes := flag.Uint("s", 5, "minutes of short break")
	longBreakMinutes := flag.Uint("l", 15, "minutes of long break")
	pomodoroAmount := flag.Uint("p", 4, "amount of total pomodoros")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	finalMessage := "You have reached the maximum number of pomodoros"

	// Cancel context on Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		finalMessage = "You have interrupted the pomodoro (Ctrl+C))"
		cancel()
	}()

	p := pomodoro.New(ctx)
	p.PomodoroMinutes = *pomodoroMinutes
	p.ShortBreakMinutes = *shortBreakMinutes
	p.LongBreakMinutes = *longBreakMinutes

	if *pomodoroAmount > math.MaxUint8 {
		*pomodoroAmount = math.MaxUint8
	}

	p.PomodoroAmount = uint8(*pomodoroAmount)

	for *pomodoroAmount*4 > uint(p.Count) {
		select {
		case <-ctx.Done():
			break
		default:
		}

		p.Start()

		if p.Count%4 == 0 {
			p.LongBreak()
		} else {
			p.ShortBreak()
		}

		if p.Count == math.MaxUint8 {
			break
		}

		p.Count++
	}

	fmt.Printf("\033[s\033[K\033[48;5;220m\033[38;5;16m %s \033[0m", finalMessage)
}
