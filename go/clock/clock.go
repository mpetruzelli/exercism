package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	hour, minute = rollover(hour, minute)
	return Clock{hour, minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, clock.minute+minutes)
}

func (clock Clock) Subtract(minutes int) Clock {
	return New(clock.hour, clock.minute-minutes)
}

func rollover(hour, minute int) (int, int) {
	if minute < 0 && minute == -60 {
	} else if minute < 0 {
		minute -= 60
	}
	hour += minute / 60
	hour = goMod(hour, 24)
	minute = goMod(minute, 60)

	return hour, minute
}

func goMod(input, divisor int) int {
	temp := input % divisor
	if temp < 0 {
		temp += divisor
	}
	return temp
}
