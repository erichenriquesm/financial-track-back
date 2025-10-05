package model

import (
	"os"
	"strings"
	"time"
)

const LayoutYYYYMMDDHHMM = "2006-01-02 15:04"

type JSONTime struct {
	time.Time
}

func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		jt.Time = time.Time{}
		return nil
	}
	loc := getAppLocation()
	t, err := time.ParseInLocation(LayoutYYYYMMDDHHMM, s, loc)
	if err != nil {
		return err
	}
	jt.Time = t.In(loc)
	return nil
}

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	if jt.Time.IsZero() {
		return []byte("null"), nil
	}
	loc := getAppLocation()
	return []byte("\"" + jt.Time.In(loc).Format(LayoutYYYYMMDDHHMM) + "\""), nil
}

func (jt JSONTime) ToTime() time.Time {
	return jt.Time.In(getAppLocation())
}

func (jt JSONTime) IsZero() bool {
	return jt.Time.IsZero()
}

func getAppLocation() *time.Location {
	tz := os.Getenv("APP_TIMEZONE")
	if tz == "" {
		tz = "America/Sao_Paulo"
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.Local
	}
	return loc
}
