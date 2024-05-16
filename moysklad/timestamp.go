package moysklad

import (
	"encoding/json"
	"time"
)

const TimestampFormat = "2006-01-02 15:04:05.000"

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an TimestampFormat.
type Timestamp struct {
	time.Time
}

func NewTimestamp(time time.Time) *Timestamp {
	return &Timestamp{Time: time}
}

func (t Timestamp) String() string {
	return t.Time.String()
}

// GetTime returns std time.Time.
func (t *Timestamp) GetTime() *time.Time {
	if t == nil {
		return nil
	}
	return &t.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format("2006-01-02 15:04:05"))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.Parse(`"`+TimestampFormat+`"`, string(data))
	return
}

// Equal reports whether t and u are equal based on time.Equal
func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time.Equal(u.Time)
}
