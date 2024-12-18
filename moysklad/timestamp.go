package moysklad

import (
	"encoding/json"
	"time"
)

// TimestampFormat Формат даты и времени.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-format-daty-i-wremeni
const TimestampFormat = "2006-01-02 15:04:05.000"

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an TimestampFormat.
type Timestamp time.Time

// NewTimestamp принимает [time.Time] и возвращает [Timestamp].
func NewTimestamp(time time.Time) *Timestamp {
	t := (Timestamp)(time)
	return &t
}

// String реализует интерфейс [fmt.Stringer].
func (timestamp Timestamp) String() string {
	return timestamp.Time().String()
}

// Time преобразует [Timestamp] в [time.Time].
func (timestamp Timestamp) Time() time.Time {
	return (time.Time)(timestamp)
}

// MarshalJSON реализует интерфейс [json.Marshaler].
func (timestamp Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal((time.Time)(timestamp).Format("2006-01-02 15:04:05"))
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (timestamp *Timestamp) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse(`"`+TimestampFormat+`"`, string(data))
	*timestamp = Timestamp(t)
	return
}
