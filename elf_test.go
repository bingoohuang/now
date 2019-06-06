package now_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/bingoohuang/now"
)

func print(n now.Now) {
	fmt.Println(n.T.Format(now.ConvertLayout("yyyy-MM-dd HH:mm:ss.SSS")), n.T.Weekday())
}

func TestParseAny(t *testing.T) {
	_, err := now.ParseAny("04 Feb 12:09")
	assert.NotNil(t, err)

	_, err = now.ParseAnyInLocation(time.Local, "04 Feb 12:09")
	assert.NotNil(t, err)

	a := now.MakeTime(time.Now().Add(-1 * time.Minute)).Format("HH:mm")
	b := now.MakeTime(time.Now().Add(1 * time.Minute)).Format("HH:mm")
	assert.True(t, now.Between(a, b))
}

func TestExample(t *testing.T) {
	print(now.BeginningOfMinute())          // 2019-06-04 10:01:00.000 Tuesday
	print(now.BeginningOfHour())            // 2019-06-04 10:00:00.000 Tuesday
	print(now.BeginningOfDay())             // 2019-06-04 00:00:00.000 Tuesday
	print(now.BeginningOfWeek(time.Monday)) // 2019-06-03 00:00:00.000 Monday
	print(now.BeginningOfWeek(time.Sunday)) // 2019-06-02 00:00:00.000 Sunday
	print(now.BeginningOfMonth())           // 2019-06-01 00:00:00.000 Saturday
	print(now.BeginningOfQuarter())         // 2019-04-01 00:00:00.000 Monday
	print(now.BeginningOfYear())            // 2019-01-01 00:00:00.000 Tuesday

	print(now.EndOfMinute())          // 2019-06-04 10:01:59.999 Tuesday
	print(now.EndOfHour())            // 2019-06-04 10:59:59.999 Tuesday
	print(now.EndOfDay())             // 2019-06-04 23:59:59.999 Tuesday
	print(now.EndOfWeek(time.Monday)) // 2019-06-09 23:59:59.999 Sunday
	print(now.EndOfWeek(time.Sunday)) // 2019-06-08 23:59:59.999 Saturday
	print(now.EndOfMonth())           // 2019-06-30 23:59:59.999 Sunday
	print(now.EndOfQuarter())         // 2019-06-30 23:59:59.999 Sunday
	print(now.EndOfYear())            // 2019-12-31 23:59:59.999 Tuesday

	// Use another time
	t1 := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.UTC)
	print(now.MakeTime(t1).EndOfMonth()) // 2013-02-28 23:59:59.999 Thursday
	print(now.Monday())                  // 2019-06-03 00:00:00.000 Monday
	print(now.Sunday())                  // 2019-06-09 00:00:00.000 Sunday
	print(now.EndOfSunday())             // 2019-06-09 23:59:59.999 Sunday
}
