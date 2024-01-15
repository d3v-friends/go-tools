package fnEnums

import (
	"fmt"
	"strings"
)

type EnumMap[T any] map[string]T

func (x *EnumMap[T]) From(v string) (enum T, err error) {
	var has bool
	if enum, has = (*x)[strings.ToLower(v)]; !has {
		err = fmt.Errorf("invalid enum string: v=%s", v)
		return
	}
	return
}
