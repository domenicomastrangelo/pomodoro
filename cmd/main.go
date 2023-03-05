package main

import (
	"flag"

	"github.com/domenicomastrangelo/pomodoro/internal/pomodoro"
)

func main() {
	pomodoroMinutes := flag.Uint("m", 25, "minutes of pomodoro")
	shortBreakMinutes := flag.Uint("s", 5, "minutes of short break")
	longBreakMinutes := flag.Uint("l", 15, "minutes of long break")

	flag.Parse()

	count := 1

	p := pomodoro.New()
	p.PomodoroMinutes = *pomodoroMinutes
	p.ShortBreakMinutes = *shortBreakMinutes
	p.LongBreakMinutes = *longBreakMinutes

	for {
		p.Start()

		if count%4 == 0 {
			p.LongBreak()
		} else {
			p.ShortBreak()
		}

		count++
	}
}
