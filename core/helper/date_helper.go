package helper

import (
	"strings"
	"time"
)

type DateHelper time.Time

var xtime *time.Time

func (c *DateHelper) UnmarshalJSON(b []byte) error {

	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}

	*c = DateHelper(t)
	xtime = &t
	return nil

}

func (c DateHelper) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}

func (c *DateHelper) Date() time.Time {
	return *xtime
}
