# Now

Now is time parser and toolkit for golang [![Go Report Card](https://goreportcard.com/badge/github.com/bingoohuang/now)](https://goreportcard.com/report/github.com/bingoohuang/now)

# Build

```bash
go get -u github.com/bingoohuang/now
```

# Usage

```go
import "github.com/bingoohuang/now"

now.Parse("13:39", "HH:mm").BeginningOfMinute()      // 2019-06-04 13:39:00.000 Tuesday

now.BeginningOfMinute()          // 2019-06-04 10:01:00.000 Tuesday
now.BeginningOfHour()            // 2019-06-04 10:00:00.000 Tuesday
now.BeginningOfDay()             // 2019-06-04 00:00:00.000 Tuesday
now.BeginningOfWeek(time.Monday)) // 2019-06-03 00:00:00.000 Monday
now.BeginningOfWeek(time.Sunday)) // 2019-06-02 00:00:00.000 Sunday
now.BeginningOfMonth()           // 2019-06-01 00:00:00.000 Saturday
now.BeginningOfQuarter()         // 2019-04-01 00:00:00.000 Monday
now.BeginningOfYear()            // 2019-01-01 00:00:00.000 Tuesday

now.EndOfMinute()          // 2019-06-04 10:01:59.999 Tuesday
now.EndOfHour()            // 2019-06-04 10:59:59.999 Tuesday
now.EndOfDay()             // 2019-06-04 23:59:59.999 Tuesday
now.EndOfWeek(time.Monday)) // 2019-06-09 23:59:59.999 Sunday
now.EndOfWeek(time.Sunday)) // 2019-06-08 23:59:59.999 Saturday
now.EndOfMonth()           // 2019-06-30 23:59:59.999 Sunday
now.EndOfQuarter()         // 2019-06-30 23:59:59.999 Sunday
now.EndOfYear()            // 2019-12-31 23:59:59.999 Tuesday

// Use another time
t1 := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.UTC)
now.MakeTime(t1).EndOfMonth() // 2013-02-28 23:59:59.999 Thursday
now.Monday()                  // 2019-06-03 00:00:00.000 Monday
now.Sunday()                  // 2019-06-09 00:00:00.000 Sunday
now.EndOfSunday()             // 2019-06-09 23:59:59.999 Sunday

```

# Thanks 

1. [now](https://github.com/jinzhu/now)
