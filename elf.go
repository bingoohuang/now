package now

import "time"

// MustParseAnyInLocation must parse string to Now in location or will panic
func MustParseAnyInLocation(loc *time.Location, strs string, formats ...string) Now {
	n := MakeTime(time.Now().In(loc))
	return n.MustParseAny(strs, formats...)
}

// MustParseAny must parse string to Now or will panic
func MustParseAny(strs string, formats ...string) Now {
	return MakeTime(time.Now()).MustParseAny(strs, formats...)
}

// BeginningOfMinute beginning of minute
func BeginningOfMinute() Now {
	return MakeTime(time.Now()).BeginningOfMinute()
}

// BeginningOfHour beginning of hour
func BeginningOfHour() Now {
	return MakeTime(time.Now()).BeginningOfHour()
}

// BeginningOfDay beginning of day
func BeginningOfDay() Now {
	return MakeTime(time.Now()).BeginningOfDay()
}

// BeginningOfWeek beginning of week
func BeginningOfWeek(weekStartDay time.Weekday) Now {
	return MakeTime(time.Now()).BeginningOfWeek(weekStartDay)
}

// BeginningOfMonth beginning of month
func BeginningOfMonth() Now {
	return MakeTime(time.Now()).BeginningOfMonth()
}

// BeginningOfQuarter beginning of quarter
func BeginningOfQuarter() Now {
	return MakeTime(time.Now()).BeginningOfQuarter()
}

// BeginningOfYear beginning of year
func BeginningOfYear() Now {
	return MakeTime(time.Now()).BeginningOfYear()
}

// EndOfMinute end of minute
func EndOfMinute() Now {
	return MakeTime(time.Now()).EndOfMinute()
}

// EndOfHour end of hour
func EndOfHour() Now {
	return MakeTime(time.Now()).EndOfHour()
}

// EndOfDay end of day
func EndOfDay() Now {
	return MakeTime(time.Now()).EndOfDay()
}

// EndOfWeek end of week
func EndOfWeek(weekStartDay time.Weekday) Now {
	return MakeTime(time.Now()).EndOfWeek(weekStartDay)
}

// EndOfMonth end of month
func EndOfMonth() Now {
	return MakeTime(time.Now()).EndOfMonth()
}

// EndOfQuarter end of quarter
func EndOfQuarter() Now {
	return MakeTime(time.Now()).EndOfQuarter()
}

// EndOfYear end of year
func EndOfYear() Now {
	return MakeTime(time.Now()).EndOfYear()
}

// Monday monday
func Monday() Now {
	return MakeTime(time.Now()).Monday()
}

// Sunday sunday
func Sunday() Now {
	return MakeTime(time.Now()).Sunday()
}

// EndOfSunday end of sunday
func EndOfSunday() Now {
	return MakeTime(time.Now()).EndOfSunday()
}

// ParseAny parse string to Now
func ParseAny(strs string, formats ...string) (Now, error) {
	n := MakeTime(time.Now())
	err := n.ParseAny(strs, formats...)

	return n, err
}

// ParseAnyInLocation parse string to Now in location
func ParseAnyInLocation(loc *time.Location, strs string, formats ...string) (Now, error) {
	n := MakeTime(time.Now().In(loc))
	err := n.ParseAny(strs, formats...)

	return n, err
}

// Between check now between the begin, end time or not
func Between(time1, time2 string) bool {
	n := MakeTime(time.Now())
	n1 := MustParseAny(time1)
	n2 := MustParseAny(time2)
	return n.Between(n1, n2)
}
