package formater

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type XTime struct {
	time.Time
}

func (t XTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`\"\"`), nil
	}
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *XTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = XTime{value}
		return nil
	}
	return fmt.Errorf("%v cannot be converted to timestamp", v)
}
