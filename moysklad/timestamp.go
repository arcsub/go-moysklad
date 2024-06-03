package moysklad

import (
	"encoding/json"
	"time"
)

const TimestampFormat = "2006-01-02 15:04:05.000"

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an TimestampFormat.
type Timestamp struct{ time.Time }

func NewTimestamp(time time.Time) *Timestamp {
	return &Timestamp{Time: time}
}

func (timestamp Timestamp) String() string {
	return timestamp.Time.String()
}

// GetTime returns std time.Time.
func (timestamp *Timestamp) GetTime() *time.Time {
	if timestamp == nil {
		return nil
	}
	return &timestamp.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (timestamp Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(timestamp.Time.Format("2006-01-02 15:04:05"))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (timestamp *Timestamp) UnmarshalJSON(data []byte) (err error) {
	timestamp.Time, err = time.Parse(`"`+TimestampFormat+`"`, string(data))
	return
}

// Equal reports whether t and u are equal based on time.Equal
func (timestamp Timestamp) Equal(u Timestamp) bool {
	return timestamp.Time.Equal(u.Time)
}
