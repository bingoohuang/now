package now_test

import (
	"testing"
	"time"

	"github.com/bingoohuang/now"
	"github.com/stretchr/testify/assert"
)

func TestBeginning(t *testing.T) {
	day := "2019-06-01 22:32:53.012"
	n, err := now.Parse(day, now.DayTimeMillisFmt)
	assert.Nil(t, err)

	assert.Equal(t, "2019-06-01 22:32:00.000", n.BeginningOfMinute().String())
	assert.Equal(t, "2019-06-01 22:00:00.000", n.BeginningOfHour().String())
	assert.Equal(t, "2019-06-01 00:00:00.000", n.BeginningOfDay().String())
	assert.Equal(t, "2019-05-27 00:00:00.000", n.BeginningOfWeek(time.Monday).String())
	assert.Equal(t, "2019-01-01 00:00:00.000", n.BeginningOfYear().String())
	assert.Equal(t, "2019-04-01 00:00:00.000", n.BeginningOfQuarter().String())
	assert.Equal(t, "2019-01-01 00:00:00.000", n.BeginningOfHalf().String())
}

func TestEnding(t *testing.T) {
	day := "2019-06-01 22:32:53.012"
	n, err := now.Parse(day, now.DayTimeMillisFmt)
	assert.Nil(t, err)

	assert.Equal(t, "2019-06-01 22:32:59.999", n.EndOfMinute().String())
	assert.Equal(t, "2019-06-01 22:59:59.999", n.EndOfHour().String())
	assert.Equal(t, "2019-06-01 23:59:59.999", n.EndOfDay().String())
	assert.Equal(t, "2019-06-30 23:59:59.999", n.EndOfMonth().String())
	assert.Equal(t, "2019-06-01 23:59:59.999", n.EndOfWeek(time.Sunday).String())
	assert.Equal(t, "2019-06-30 23:59:59.999", n.EndOfQuarter().String())
	assert.Equal(t, "2019-06-30 23:59:59.999", n.EndOfHalf().String())
	assert.Equal(t, "2019-12-31 23:59:59.999", n.EndOfYear().String())
}

func between(n now.Now, a, b string) bool {
	an, _ := now.Parse(a, "HH:mm")
	bn, _ := now.Parse(b, "HH:mm")
	nb := n.BeginningOfMinute()
	ab := an.BeginningOfMinute()
	bb := bn.BeginningOfMinute()
	return nb.Between(ab, bb)
}

func TestWithinWorkingTime(t *testing.T) {
	a := assert.New(t)
	n, err := now.Parse("13:39", "HH:mm")
	assert.Nil(t, err)

	a.True(between(n, "13:39", "13:49"), "时间在区间之内")
	a.True(between(n, "13:39", "13:49"), "时间在区间之内")
	a.True(between(n, "13:29", "13:39"), "时间在区间之内")
	a.False(between(n, "13:40", "13:49"), "时间在区间之内")
	a.False(between(n, "13:28", "13:38"), "时间在区间之内")
}

func RoundMinutes(t now.Now, offset time.Duration) now.Now {
	return t.BeginningOfMinute().Offset(offset)
}

func TestRoundMinutes(t *testing.T) {
	a := assert.New(t)
	t1, _ := now.Parse("13:39:30", "HH:mm:ss")

	t10, _ := now.Parse("13:39:00", "HH:mm:ss")
	t11, _ := now.Parse("13:38:00", "HH:mm:ss")
	t12, _ := now.Parse("13:40:00", "HH:mm:ss")
	a.Equal(RoundMinutes(t1, time.Duration(0)), t10.BeginningOfMinute(), "规整到当前整分")
	a.Equal(RoundMinutes(t1, time.Duration(-1)*time.Minute), t11.BeginningOfMinute(), "规整到-1整分")
	a.Equal(RoundMinutes(t1, time.Duration(1)*time.Minute), t12.BeginningOfMinute(), "规整到+1整分")

	du, err := time.ParseDuration("-3m")
	a.Nil(err)
	a.Equal(du, time.Duration(-3)*time.Minute)
}
