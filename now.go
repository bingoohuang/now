package now

import "time"

// Now now struct
type Now struct {
	time.Time
	// WeekStartDay set week start day, default is sunday
	WeekStartDay time.Weekday
}

// New initialize Now with time
func New(t time.Time) *Now {
	return &Now{Time: t, WeekStartDay: time.Monday}
}

// NewNow initialize Now with now time
func NewNow() *Now {
	return New(time.Now())
}
