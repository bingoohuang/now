package now

import (
	"strings"
	"time"
)

// Format formats t in layout.
func (n Now) Format(layout string) string {
	return n.T.Format(ConvertLayout(layout))
}

// Parse parses d in layout to now
func Parse(d, layout string) (now Now, err error) {
	now = MakeNow()
	err = now.Parse(d, layout)
	return now, err
}

// Parse parses d in layout to time.Time
func (n *Now) Parse(d, layout string) (err error) {
	lo := ConvertLayout(layout)
	var t time.Time
	t, err = time.ParseInLocation(lo, d, time.Local)
	if err != nil {
		return err
	}

	n.fill(layout, t)
	return nil
}

func (n *Now) fill(layout string, fill time.Time) {
	year := n.T.Year()
	if strings.Contains(layout, "yy") {
		year = fill.Year()
	}
	month := n.T.Month()
	if strings.Contains(layout, "MM") {
		month = fill.Month()
	}
	day := n.T.Day()
	if strings.Contains(layout, "dd") {
		day = fill.Day()
	}
	hour := n.T.Hour()
	if strings.Contains(layout, "HH") {
		hour = fill.Hour()
	}
	minute := n.T.Minute()
	if strings.Contains(layout, "mm") {
		minute = fill.Minute()
	}
	second := n.T.Second()
	if strings.Contains(layout, "ss") {
		second = fill.Second()
	}
	nano := n.T.Nanosecond()
	if strings.Contains(layout, "SSS") {
		nano = fill.Nanosecond()
	}
	n.T = time.Date(year, month, day, hour, minute, second, nano, time.Local)
	n.present()
}
