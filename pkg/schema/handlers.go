package schema

import (
	"encoding/json"
	"fmt"
)

type ExpiredBool bool
type ExpiredString string

type ExpiredType struct {
	Data interface{}
}

func (t *ExpiredType) UnmarshalJSON(data []byte) error {
	var a ExpiredBool
	err := json.Unmarshal(data, &a)
	if err == nil {
		t.Data = a
		return nil
	}

	if _, ok := err.(*json.UnmarshalTypeError); err != nil && !ok {
		return err
	}

	var b ExpiredString
	err = json.Unmarshal(data, &b)
	if err != nil {
		return err
	}
	t.Data = b
	return nil
}

func (t *ExpiredType) String() string {
	switch d := t.Data.(type) {
	case ExpiredBool:
		return fmt.Sprint(d)
	case ExpiredString:
		return fmt.Sprintf("%s", d)
	}
	return ""
}
