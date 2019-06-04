package now_test

import (
	"encoding/json"
	"testing"

	"github.com/bingoohuang/now"
	"github.com/stretchr/testify/assert"
)

func TestNowJSON(t *testing.T) {
	type MyBean struct {
		Now now.Now `json:"now"`
	}

	b := MyBean{Now: now.MakeNow().RoundNano()}
	bytes, err := json.Marshal(b)
	assert.Nil(t, err)

	var c MyBean
	err = json.Unmarshal(bytes, &c)
	assert.Nil(t, err)

	assert.Equal(t, b, c, "equals")
}
