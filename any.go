package now

import (
	"errors"
	"regexp"
	"time"
)

// Formats default time formats will be parsed as
var Formats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4",
	"2006-1-2 15:4:5", "1-2", "15:4:5", "15:4", "15", "15:4:5 Jan 2, 2006 MST",
	"2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05-07:00"}

func parseWithFormat(str string) (t time.Time, err error) {
	for _, fmt := range Formats {
		if t, err = time.Parse(fmt, str); err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, etc
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)      // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc

// ParseAny parse string to time
func (n *Now) ParseAny(strs ...string) (err error) {
	var setCurrentTime bool
	var parseTime []int

	x := n.T
	var currentTime = []int{x.Nanosecond(), x.Second(), x.Minute(), x.Hour(), x.Day(), int(x.Month()), x.Year()}
	var currentLocation = x.Location()
	var onlyTimeInStr = true
	var t time.Time

	for _, str := range strs {
		hasTimeInStr := hasTimeRegexp.MatchString(str) // match 15:04:05, 15
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)

		if t, err = parseWithFormat(str); err == nil {
			location := t.Location()
			if location.String() == "UTC" {
				location = currentLocation
			}

			parseTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = []int{t.Nanosecond(), t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
		}
	}

	n.T = t
	n.present()
	return err
}

// MustParseAny must parse string to time or it will panic
func (n Now) MustParseAny(strs ...string) Now {
	if err := n.ParseAny(strs...); err != nil {
		panic(err)
	}
	return n
}
