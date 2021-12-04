package domain

import (
	"fmt"
	"strings"
	"time"
)

type customTime time.Time

const CtLayout = time.RFC3339
const CtLayoutNoTime = "2006-01-02"

func (ct *customTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	arr := strings.Split(s, " ")
	var nt time.Time
	if len(arr) > 1 {
		nt, err = time.Parse(CtLayout, s)
	} else {
		nt, err = time.Parse(CtLayoutNoTime, s)
	}
	*ct = customTime(nt)
	return
}

func (ct *customTime) Scan(v interface{}) error {
	vt, err := time.Parse(CtLayout, string(v.([]byte)))
	if err != nil {
		return err
	}
	*ct = customTime(vt)
	return nil
}

func (ct customTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *customTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf(t.Format(CtLayout))
}

func (ct *customTime) Time() time.Time {
	return time.Time(*ct)
}