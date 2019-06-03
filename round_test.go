package now_test

import (
	"testing"

	"github.com/bingoohuang/now"
	"github.com/stretchr/testify/assert"
)

func TestNow_BeginningOfDay(t *testing.T) {
	day := "2019-06-01 22:32:53.012"
	n := now.NewNow()
	err := n.Parse(day, now.DayTimeMillisFmt)
	assert.Nil(t, err)

	assert.Equal(t, "2019-06-01 22:32:00.000", n.BeginningOfMinute().Format(now.DayTimeMillisFmt))
	assert.Equal(t, "2019-06-01 22:00:00.000", n.BeginningOfHour().Format(now.DayTimeMillisFmt))
	assert.Equal(t, "2019-06-01 00:00:00.000", n.BeginningOfDay().Format(now.DayTimeMillisFmt))
	assert.Equal(t, "2019-05-27 00:00:00.000", n.BeginningOfWeek().Format(now.DayTimeMillisFmt))
	assert.Equal(t, "2019-01-01 00:00:00.000", n.BeginningOfYear().Format(now.DayTimeMillisFmt))
}
