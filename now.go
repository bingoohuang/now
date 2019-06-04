package now

import "time"

// Now now struct
type Now struct {
	// T time
	T time.Time
	// P presentation
	P string
	// L layout
	L string
}

func (n *Now) present() {
	n.P = n.T.Format(n.L)
}

// Make initialize Now with time
func Make(t time.Time, layout string) Now {
	n := Now{T: t, L: ConvertLayout(layout)}
	n.present()

	return n
}

// MakeTime initialize Now with time
func MakeTime(t time.Time) Now {
	return Make(t, DayTimeMillisFmt)
}

// MakeNow initialize Now with now time
func MakeNow() Now {
	return MakeTime(time.Now())
}

// String returns n's presentation.
func (n Now) String() string {
	return n.P
}

// Between tells whether n between a and b.
func (n Now) Between(a, b Now) bool {
	t := n.T
	at := a.T
	bt := b.T

	return t.Equal(at) || t.After(at) && t.Before(bt) || t.Equal(bt)
}
