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

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return
	}
	var str string
	err = json.Unmarshal(b, &str)
	if err != nil {
		return
	}
	r, err := time.Parse(time.RFC3339, str)
	if err != nil {
		r, err = time.Parse(iso8601Layout, str)
		if err != nil {
			return
		}
	}
	t.Time = r
	return nil
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return
	}
	var str string
	err = json.Unmarshal(b, &str)
	if err != nil {
		return
	}
	r, err := time.Parse(time.RFC3339, str)
	if err != nil {
		r, err = time.Parse(time.DateOnly, str)
		if err != nil {
			return
		}
	}
	d.Time = r
	return nil
}
