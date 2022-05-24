package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	hour := h + m/60
	if m%60 < 0 {
		hour--
	}
	hour %= 24
	if hour < 0 {
		hour = 24 + hour
	}
	minute := m % 60
	if minute < 0 {
		minute = 60 + minute
	}
	return Clock{
		hour:   hour,
		minute: minute,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
