package fnCase

import (
	"github.com/d3v-friends/go-tools/fn/fnReflect"
	"github.com/gertd/go-pluralize"
)

type String string

var client = pluralize.NewClient()

func NewString(v string) *String {
	return fnReflect.Pointer(String(v))
}

func (x *String) String() string {
	return string(*x)
}

func (x *String) PascalCase() *String {
	return fnReflect.Pointer(String(PascalCase(x.String())))
}

func (x *String) CamelCase() *String {
	return fnReflect.Pointer(String(CamelCase(x.String())))
}

func (x *String) SnakeCase() *String {
	return fnReflect.Pointer(String(SnakeCase(x.String())))
}

func (x *String) Pluralize() *String {
	return fnReflect.Pointer(String(client.Plural(x.String())))
}

func (x *String) Singular() *String {
	return fnReflect.Pointer(String(client.Singular(x.String())))
}
