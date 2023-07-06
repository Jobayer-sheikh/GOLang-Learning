package database

import (
	"database/sql/driver"
	"time"
)

type NullString string
type NullTime time.Time

const (
	emptyString = ""
)

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = emptyString
	} else {
		*s = NullString(value.(string))
	}

	return nil
}

func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	return string(s), nil
}

func (t *NullTime) Scan(value interface{}) error {
	if t != nil {
		*t = NullTime(value.(time.Time))
	}

	return nil
}

func (t *NullTime) Value() (driver.Value, error) {
	return time.Time(*t), nil
}
