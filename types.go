package sfdc

import (
	"encoding/json"
	"time"
)

type Time struct {
	time.Time
}

type Date struct {
	time.Time
}

const iso8601Layout string = "2006-01-02T15:04:05Z0700"

func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	if str == "null" || str == "" {
		return nil
	}
	r, err := time.Parse(time.RFC3339, str)
	if err != nil {
		r, err = time.Parse(iso8601Layout, str)
		if err != nil {
			return err
		}
	}
	t.Time = r
	return nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	if str == "null" || str == "" {
		return nil
	}
	r, err := time.Parse(time.RFC3339, str)
	if err != nil {
		r, err = time.Parse(time.DateOnly, str)
		if err != nil {
			return err
		}
	}
	d.Time = r
	return nil
}
