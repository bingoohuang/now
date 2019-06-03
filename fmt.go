package now

import (
	"strings"
	"time"
)

const DayFmt = "yyyy-MM-dd"
const TimeFmt = "HH:mm:ss"
const DayTimeFmt = DayFmt + " " + TimeFmt
const DayTimeMillisFmt = DayTimeFmt + ".SSS"

// Format formats t in layout.
func (now Now) Format(layout string) string {
	lo := parseLayout(layout)
	return now.Time.Format(lo)
}

func parseLayout(layout string) string {
	lo := layout
	lo = strings.Replace(lo, "yyyy", "2006", -1)
	lo = strings.Replace(lo, "yy", "06", -1)
	lo = strings.Replace(lo, "MM", "01", -1)
	lo = strings.Replace(lo, "dd", "02", -1)
	lo = strings.Replace(lo, "HH", "15", -1)
	lo = strings.Replace(lo, "mm", "04", -1)
	lo = strings.Replace(lo, "ss", "05", -1)
	lo = strings.Replace(lo, "SSS", "000", -1)
	return lo
}

// Parse parses d in layout to time.Time
func (now *Now) Parse(d, layout string) (err error) {
	lo := parseLayout(layout)
	var t time.Time
	t, err = time.ParseInLocation(lo, d, time.Local)
	if err != nil {
		return err
	}

	now.fill(layout, t)
	return nil
}

func (now *Now) fill(layout string, fill time.Time) {
	year := now.Time.Year()
	if strings.Contains(layout, "yy") {
		year = fill.Year()
	}
	month := now.Time.Month()
	if strings.Contains(layout, "MM") {
		month = fill.Month()
	}
	day := now.Time.Day()
	if strings.Contains(layout, "dd") {
		day = fill.Day()
	}
	hour := now.Time.Hour()
	if strings.Contains(layout, "HH") {
		hour = fill.Hour()
	}
	minute := now.Time.Minute()
	if strings.Contains(layout, "mm") {
		minute = fill.Minute()
	}
	second := now.Time.Second()
	if strings.Contains(layout, "ss") {
		second = fill.Second()
	}
	nano := now.Time.Nanosecond()
	if strings.Contains(layout, "SSS") {
		nano = fill.Nanosecond()
	}
	now.Time = time.Date(year, month, day, hour, minute, second, nano, time.Local)
}
