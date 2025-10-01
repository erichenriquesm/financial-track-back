package utils

import (
	"time"
)

type JSONTime time.Time

// MarshalJSON formata o tempo no estilo "2006-01-02 15:04:05"
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).Format("2006-01-02 15:04:05")
	return []byte(`"` + formatted + `"`), nil
}

// Se precisar, também pode implementar UnmarshalJSON
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	*t = JSONTime(parsed)
	return nil
}
