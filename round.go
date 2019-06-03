package now

import "time"

// BeginningOfMinute beginning of minute
func (now *Now) BeginningOfMinute() *Now {
	y, m, d := now.Date()
	now.Time = time.Date(y, m, d, now.Time.Hour(), now.Time.Minute(), 0, 0, now.Time.Location())
	return now
}

// BeginningOfMinute beginning of hour
func (now *Now) BeginningOfHour() *Now {
	y, m, d := now.Date()
	now.Time = time.Date(y, m, d, now.Time.Hour(), 0, 0, 0, now.Time.Location())
	return now
}

// BeginningOfDay beginning of day
func (now *Now) BeginningOfDay() *Now {
	y, m, d := now.Date()
	now.Time = time.Date(y, m, d, 0, 0, 0, 0, now.Time.Location())
	return now
}

// BeginningOfWeek beginning of week
func (now *Now) BeginningOfWeek() *Now {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())

	if now.WeekStartDay != time.Sunday {
		weekStartDayInt := int(now.WeekStartDay)

		if weekday < weekStartDayInt {
			weekday += 7 - weekStartDayInt
		} else {
			weekday -= weekStartDayInt
		}
	}
	now.Time = t.AddDate(0, 0, -weekday)
	return now
}

// BeginningOfMonth beginning of month
func (now *Now) BeginningOfMonth() *Now {
	y, m, _ := now.Date()
	now.Time = time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
	return now
}

// BeginningOfQuarter beginning of quarter
func (now *Now) BeginningOfQuarter() *Now {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	now.Time = month.AddDate(0, -offset, 0)
	return now
}

// BeginningOfHalf beginning of half year
func (now *Now) BeginningOfHalf() *Now {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	now.Time = month.AddDate(0, -offset, 0)
	return now
}

// BeginningOfYear BeginningOfYear beginning of year
func (now *Now) BeginningOfYear() *Now {
	y, _, _ := now.Date()
	now.Time = time.Date(y, time.January, 1, 0, 0, 0, 0, now.Location())
	return now
}

// EndOfMinute end of minute
func (now *Now) EndOfMinute() *Now {
	now.Time = now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
	return now
}

// EndOfHour end of hour
func (now *Now) EndOfHour() *Now {
	now.Time = now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
	return now
}

// EndOfDay end of day
func (now *Now) EndOfDay() *Now {
	y, m, d := now.Date()
	now.Time = time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
	return now
}

// EndOfWeek end of week
func (now *Now) EndOfWeek() *Now {
	now.Time = now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
	return now
}

// EndOfMonth end of month
func (now *Now) EndOfMonth() *Now {
	now.Time = now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
	return now
}

// EndOfQuarter end of quarter
func (now *Now) EndOfQuarter() *Now {
	now.Time = now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
	return now
}

// EndOfHalf end of half year
func (now *Now) EndOfHalf() *Now {
	now.Time = now.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
	return now
}

// EndOfYear end of year
func (now *Now) EndOfYear() *Now {
	now.Time = now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
	return now
}

// Monday monday
func (now *Now) Monday() *Now {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	now.Time = t.AddDate(0, 0, -weekday+1)
	return now
}

// Sunday sunday
func (now *Now) Sunday() *Now {
	now.BeginningOfDay()
	weekday := int(now.Weekday())
	if weekday != 0 {
		now.Time = now.AddDate(0, 0, (7 - weekday))
	}
	return now
}

// EndOfSunday end of sunday
func (now *Now) EndOfSunday() *Now {
	New(now.Sunday().Time).EndOfDay()
	return now
}
