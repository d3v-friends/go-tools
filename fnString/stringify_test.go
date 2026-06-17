package fnString_test

import (
	"fmt"
	"testing"

	"github.com/d3v-friends/go-tools/fnString"
	"github.com/stretchr/testify/assert"
)

func TestStringify(test *testing.T) {
	test.Run("stringify", func(t *testing.T) {
		var str = fnString.Stringify(map[string]any{
			"test1": 1,
			"test2": "string",
			"test3": []string{"a", "b"},
			"test4": map[string]any{"a": 1, "b": "string"},
		})

		assert.Equal(t, "test1=1, test2=string, test3=[a b], test4=map[a:1 b:string]", str)
		fmt.Printf("%s\n", str)
	})
}
