package now

import (
	"strings"
)

// MarshalJSON marshal n to JSON
func (n Now) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(n.L)+2)
	b = append(b, '"')
	b = n.T.AppendFormat(b, n.L)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON unmarshal JSON buf to n
func (n *Now) UnmarshalJSON(buf []byte) error {
	s := strings.Trim(string(buf), `"`)
	var err error
	*n, err = Parse(s, DayTimeMillisFmt)

	return err
}
