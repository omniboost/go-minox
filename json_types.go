package minox

import (
	"encoding/json"
	"strconv"
	"time"
)

type Bool bool

func (b Bool) MarshalJSON() ([]byte, error) {
	if b == true {
		return json.Marshal("1")
	}
	return json.Marshal("0")
}

type Int int

func (i *Int) UnmarshalJSON(data []byte) error {
	realInt := 0
	err := json.Unmarshal(data, &realInt)
	if err == nil {
		*i = Int(realInt)
		return nil
	}

	// error, so maybe it isn't an int
	str := ""
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if str == "" {
		*i = 0
		return nil
	}

	realInt, err = strconv.Atoi(str)
	if err != nil {
		return err
	}

	i2 := Int(realInt)
	*i = i2
	return nil
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	return err
}
