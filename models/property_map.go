package models

import (
	"database/sql/driver"
	"encoding/json"
)

type PropertyMap map[string]interface{}

func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *PropertyMap) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		source, _ := src.([]byte)
		var i interface{}
		if err := json.Unmarshal(source, &i); err != nil {
			return err
		}

		*p, _ = i.(map[string]interface{})
	}
	return nil
}
