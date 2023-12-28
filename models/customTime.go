package models

import (
	"database/sql/driver"
	"strings"
	"time"
)

const customTimeLayout = "2006-01-02"

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	c.Time, err = time.Parse(customTimeLayout, s)
	return
}

func (c *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + c.Format(customTimeLayout) + "\""), nil
}

func (c CustomTime) Value() (driver.Value, error) {
	return c.Time, nil
}

func (c *CustomTime) String() string {
	return c.Format(customTimeLayout)
}

func (c *CustomTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		c.Time = v
	case []byte:
		s := string(v)
		t, err := time.Parse(customTimeLayout, s)
		if err != nil {
			return err
		}
		c.Time = t
	case string:
		t, err := time.Parse(customTimeLayout, v)
		if err != nil {
			return err
		}
		c.Time = t
	default:
		return nil
	}
	return nil
}
