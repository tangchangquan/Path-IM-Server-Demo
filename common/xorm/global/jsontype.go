package global

import (
	"database/sql/driver"
	"encoding/json"
)

type Json []byte
type SliceString []string
type MapStringBool map[string]bool

func (o Json) Value() (driver.Value, error) {
	return json.Marshal(o)
}
func (o *Json) Scan(input interface{}) error {
	s := string(input.([]byte))
	err := json.Unmarshal([]byte(s), o)
	if err != nil {

	}
	return nil
}
func (o SliceString) Value() (driver.Value, error) {
	return json.Marshal(o)
}
func (o *SliceString) Scan(input interface{}) error {
	s := string(input.([]byte))
	err := json.Unmarshal([]byte(s), o)
	if err != nil {

	}
	return nil
}
func (o MapStringBool) Value() (driver.Value, error) {
	return json.Marshal(o)
}
func (o *MapStringBool) Scan(input interface{}) error {
	s := string(input.([]byte))
	err := json.Unmarshal([]byte(s), o)
	if err != nil {

	}
	return nil
}
