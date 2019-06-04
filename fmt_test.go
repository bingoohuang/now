package now_test

import (
	"testing"

	"github.com/bingoohuang/now"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	day := "2019-06-03 22:32:53.012"
	layout := now.DayTimeMillisFmt
	n := now.MakeNow()
	err := n.Parse(day, layout)
	assert.Nil(t, err)

	day2 := n.Format(layout)
	assert.Equal(t, day2, day)
}

func TestHHMM(t *testing.T) {
	day := "2019-06-01 22:32:53.012"
	layout := now.DayTimeMillisFmt
	n := now.MakeNow()
	err := n.Parse(day, layout)
	assert.Nil(t, err)

	err = n.Parse("22:00:00.000", "HH:mm:ss.SSS")
	assert.Nil(t, err)

	assert.Equal(t, n.String(), "2019-06-01 22:00:00.000")
}
