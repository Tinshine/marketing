package gormtype

import (
	"database/sql/driver"
	"marketing/consts/errs"

	"github.com/bytedance/sonic"
)

type Slice[T any] []T

func (d *Slice[T]) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return sonic.Unmarshal(v, d)
	case string:
		return sonic.UnmarshalString(v, d)
	default:
		return errs.GormScanWrongType
	}
}

func (d *Slice[T]) Value() (driver.Value, error) {
	return sonic.MarshalString(d)
}
