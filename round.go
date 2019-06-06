package now

import "time"

// BeginningOfMinute beginning of minute
func (n Now) BeginningOfMinute() Now {
	y, m, d := n.T.Date()
	n.T = time.Date(y, m, d, n.T.Hour(), n.T.Minute(), 0, 0, n.T.Location())
	n.present()
	return n
}

// BeginningOfMinute beginning of hour
func (n Now) BeginningOfHour() Now {
	y, m, d := n.T.Date()
	n.T = time.Date(y, m, d, n.T.Hour(), 0, 0, 0, n.T.Location())
	n.present()
	return n
}

// BeginningOfDay beginning of day
func (n Now) BeginningOfDay() Now {
	y, m, d := n.T.Date()
	n.T = time.Date(y, m, d, 0, 0, 0, 0, n.T.Location())
	n.present()
	return n
}

// BeginningOfWeek beginning of week
func (n Now) BeginningOfWeek(weekStartDay time.Weekday) Now {
	t := n.BeginningOfDay()
	weekday := int(t.T.Weekday()) - int(weekStartDay)

	n.T = t.T.AddDate(0, 0, -weekday)
	n.present()
	return n
}

// BeginningOfMonth beginning of month
func (n Now) BeginningOfMonth() Now {
	y, m, _ := n.T.Date()
	n.T = time.Date(y, m, 1, 0, 0, 0, 0, n.T.Location())
	n.present()
	return n
}

// BeginningOfQuarter beginning of quarter
func (n Now) BeginningOfQuarter() Now {
	month := n.BeginningOfMonth()
	offset := (int(month.T.Month()) - 1) % 3
	n.T = month.T.AddDate(0, -offset, 0)
	n.present()
	return n
}

// BeginningOfHalf beginning of half year
func (n Now) BeginningOfHalf() Now {
	month := n.BeginningOfMonth()
	offset := (int(month.T.Month()) - 1) % 6
	n.T = month.T.AddDate(0, -offset, 0)
	n.present()
	return n
}

// BeginningOfYear BeginningOfYear beginning of year
func (n Now) BeginningOfYear() Now {
	y, _, _ := n.T.Date()
	n.T = time.Date(y, time.January, 1, 0, 0, 0, 0, n.T.Location())
	n.present()
	return n
}

// EndOfMinute end of minute
func (n Now) EndOfMinute() Now {
	n.T = n.BeginningOfMinute().T.Add(time.Minute - time.Nanosecond)
	n.present()
	return n
}

// EndOfHour end of hour
func (n Now) EndOfHour() Now {
	n.T = n.BeginningOfHour().T.Add(time.Hour - time.Nanosecond)
	n.present()
	return n
}

// EndOfDay end of day
func (n Now) EndOfDay() Now {
	y, m, d := n.T.Date()
	n.T = time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), n.T.Location())
	n.present()
	return n
}

// EndOfWeek end of week
func (n Now) EndOfWeek(weekStartDay time.Weekday) Now {
	t := n.BeginningOfWeek(weekStartDay)
	n.T = t.T.AddDate(0, 0, 7).Add(-time.Nanosecond)
	n.present()
	return n
}

// EndOfMonth end of month
func (n Now) EndOfMonth() Now {
	n.T = n.BeginningOfMonth().T.AddDate(0, 1, 0).Add(-time.Nanosecond)
	n.present()
	return n
}

// EndOfQuarter end of quarter
func (n Now) EndOfQuarter() Now {
	n.T = n.BeginningOfQuarter().T.AddDate(0, 3, 0).Add(-time.Nanosecond)
	n.present()
	return n
}

// EndOfHalf end of half year
func (n Now) EndOfHalf() Now {
	n.T = n.BeginningOfHalf().T.AddDate(0, 6, 0).Add(-time.Nanosecond)
	n.present()
	return n
}

// EndOfYear end of year
func (n Now) EndOfYear() Now {
	n.T = n.BeginningOfYear().T.AddDate(1, 0, 0).Add(-time.Nanosecond)
	n.present()
	return n
}

// Monday monday
func (n Now) Monday() Now {
	t := n.BeginningOfDay()
	weekday := int(t.T.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	n.T = t.T.AddDate(0, 0, -weekday+1)
	n.present()
	return n
}

// Sunday sunday
func (n Now) Sunday() Now {
	n.BeginningOfDay()
	weekday := int(n.T.Weekday())
	if weekday != 0 {
		n.T = n.T.AddDate(0, 0, 7-weekday)
	}
	return n.BeginningOfDay()
}

// EndOfSunday end of sunday
func (n Now) EndOfSunday() Now {
	return n.Sunday().EndOfDay()
}

// Offset add offset to n
func (n Now) Offset(offset time.Duration) Now {
	n.T = n.T.Add(offset)
	n.present()
	return n
}

// RoundNano add offset to n
func (n Now) RoundNano() Now {
	y, m, d := n.T.Date()
	n.T = time.Date(y, m, d, n.T.Hour(), n.T.Minute(), n.T.Second(), 0, n.T.Location())
	n.present()
	return n
}
