package moysklad

import (
	"encoding/json"
	"time"
)

// TimestampFormat Формат даты и времени
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-format-daty-i-wremeni
const TimestampFormat = "2006-01-02 15:04:05.000"

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an TimestampFormat.
type Timestamp time.Time

func NewTimestamp(time time.Time) *Timestamp {
	t := (Timestamp)(time)
	return &t
}

func (timestamp Timestamp) String() string {
	return timestamp.Time().String()
}

func (timestamp Timestamp) Time() time.Time {
	return (time.Time)(timestamp)
}

// MarshalJSON implements the json.Marshaler interface.
func (timestamp Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal((time.Time)(timestamp).Format("2006-01-02 15:04:05"))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (timestamp *Timestamp) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse(`"`+TimestampFormat+`"`, string(data))
	*timestamp = Timestamp(t)
	return
}
