package domain

import (
	"fmt"
	"strings"
	"time"
)

type customTime time.Time

const ctLayout = time.RFC3339

func (ct *customTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = customTime(nt)
	return
}

func (ct customTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *customTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf(t.Format(ctLayout))
}

func (ct *customTime) Time() time.Time {
	return time.Time(*ct)
}