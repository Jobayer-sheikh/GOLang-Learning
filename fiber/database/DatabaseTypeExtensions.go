package database

import (
	"fmt"
	"time"
)

const (
	dateLayout = "2006-01-02 15:04:05 Z07:00"
)

func (t *NullTime) UnmarshalJSON(b []byte) error {
	var err error
	var tm time.Time

	if tm, err = time.Parse(dateLayout, string(b)); err == nil {
		*t = NullTime(tm)
	}

	return err
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *NullTime) String() string {
	tm := time.Time(*t)
	return fmt.Sprintf("%q", tm.Format(dateLayout))
}
