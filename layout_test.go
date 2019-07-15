package now

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatTime(t *testing.T) {
	t1 := time.Date(2013, 02, 18, 17, 51, 49, 123000000, time.Local)
	f := FormatTime(t1, "yyyy-MM-dd HH:mm:ss.SSS")
	assert.Equal(t, "2013-02-18 17:51:49.123", f)

	t2, err := ParseTimeLocal(f, "yyyy-MM-dd HH:mm:ss.SSS")
	assert.Nil(t, err)
	assert.Equal(t, t1, t2)
}
