package fnGoroutine_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/d3v-friends/go-tools/fnGoroutine"
)

func TestDivide(test *testing.T) {
	test.Run("divide", func(t *testing.T) {
		fnGoroutine.DivideList(
			context.TODO(),
			100,
			20,
			func(ctx context.Context, page int, size int, total int) (err error) {
				fmt.Printf("page: %d, size: %d, total: %d\n", page, size, total)
				return
			},
		)
	})
}
