package fnGoroutine

import (
	"context"

	"github.com/d3v-friends/go-tools/fnDefault"
)

type FnDivideGoroutine func(ctx context.Context, page int, size int, total int) (err error)

func DivideList(
	ctx context.Context,
	total int,
	unit int,
	fn FnDivideGoroutine,
	thread ...int,
) []error {

	var page = total / unit
	if total%unit != 0 {
		page += 1
	}

	var jobs = make([]Job, page)
	for i := 0; i < page; i++ {
		jobs[i] = func(ctx context.Context) (err error) {
			return fn(ctx, i, unit, total)
		}
	}

	var worker = NewThread(fnDefault.Param(thread, 10))
	return worker.Do(ctx, jobs)
}
