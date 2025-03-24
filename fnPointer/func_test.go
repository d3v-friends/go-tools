package fnPointer_test

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Test struct {
	Name string
}

func (x Test) String() string {
	return x.Name
}

func TestReflect(test *testing.T) {
	test.Run("ptr", func(t *testing.T) {
		var i = "abcd"

		assert.Equal(t, reflect.String, reflect.TypeOf(i).Kind())
		assert.Equal(t, reflect.Pointer, reflect.TypeOf(&i).Kind())
		assert.Equal(t, reflect.Ptr, reflect.TypeOf(&i).Kind())

		var l = []string{"a"}
		assert.Equal(t, reflect.Slice, reflect.TypeOf(l).Kind())
		assert.Equal(t, reflect.Pointer, reflect.TypeOf(&l).Kind())

		var s = Test{
			Name: "test",
		}

		assert.Equal(t, reflect.Struct, reflect.TypeOf(s).Kind())
		assert.Equal(t, reflect.Pointer, reflect.TypeOf(&s).Kind())

	})
}
