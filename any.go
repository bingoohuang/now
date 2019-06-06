package now

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// Formats default time formats will be parsed as
const Formats = "1/2/2006^" +
	"1/2/2006 15:4:5^" +
	"2006^" +
	"2006-1^" +
	"2006-1-2^" +
	"2006-1-2 15^" +
	"2006-1-2 15:4^" +
	"2006-1-2 15:4:5^" +
	"1-2^" +
	"15:4:5^" +
	"15:4^" +
	"15^" +
	"15:4:5 Jan 2, 2006 MST^" +
	"2006-01-02 15:04:05.999999999 -0700 MST^" +
	"2006-01-02T15:04:05-07:00"

func parseWithFormat(str string, formats []string) (t time.Time, err error) {
	for _, fmt := range formats {
		if t, err = time.Parse(fmt, str); err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

// ParseAny parse string to time
func (n *Now) ParseAny(strs string, formats ...string) (err error) {
	// match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, etc
	var hasTimeReg = regexp.MustCompile(`(\s+|^\s*)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)
	// match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc
	var onlyTimeReg = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)

	var setCurrentTime bool
	var pt []int

	x := n.T
	var currentTime = []int{x.Nanosecond(), x.Second(), x.Minute(), x.Hour(), x.Day(), int(x.Month()), x.Year()}
	var currentLocation = x.Location()
	var onlyTime = true
	var t time.Time

	parts := strings.Split(strs, "^")
	fmts := strings.Split(Formats+"^"+strings.Join(formats, "^"), "^")

	for _, str := range parts {
		hasTime := hasTimeReg.MatchString(str) // match 15:04:05, 15
		onlyTime = hasTime && onlyTime && onlyTimeReg.MatchString(str)

		t, err = parseWithFormat(str, fmts)
		if err != nil {
			continue
		}

		location := t.Location()
		if location.String() == "UTC" {
			location = currentLocation
		}

		pt = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}

		for i, v := range pt {
			// Don't reset hour, minute, second if current time str including time
			if hasTime && i <= 3 {
				continue
			}

			// If value is zero, replace it with current time
			if v == 0 {
				if setCurrentTime {
					pt[i] = currentTime[i]
				}
			} else {
				setCurrentTime = true
			}

			// if current time only includes time, should change day, month to current time
			if onlyTime {
				if i == 4 || i == 5 {
					pt[i] = currentTime[i]
					continue
				}
			}
		}

		t = time.Date(pt[6], time.Month(pt[5]), pt[4], pt[3], pt[2], pt[1], pt[0], location)
		currentTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
	}

	n.T = t
	n.present()
	return err
}

// MustParseAny must parse string to time or it will panic
func (n Now) MustParseAny(strs string, formats ...string) Now {
	if err := n.ParseAny(strs, formats...); err != nil {
		panic(err)
	}
	return n
}
