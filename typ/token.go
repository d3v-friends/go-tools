package typ

import (
	"database/sql/driver"
	"fmt"
)

type Token string

func (x *Token) String() string {
	return string(*x)
}

func (x *Token) Scan(src any) (err error) {
	switch v := src.(type) {
	case string:
		*x = Token(v)
		return
	case []byte:
		*x = Token(v)
		return
	default:
		err = fmt.Errorf("invalid src type")
		return
	}
}

func (x Token) Value() (res driver.Value, err error) {
	res = string(x)
	return
}
