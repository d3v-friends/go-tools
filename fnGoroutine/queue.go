package fnGoroutine

import (
	"context"

	"github.com/d3v-friends/go-tools/fnLogger"
)

type Queue struct {
	ch chan Job
	cg ContextGenerator
}

func NewQueue(
	cg ContextGenerator,

) *Queue {
	return &Queue{
		ch: make(chan Job),
		cg: cg,
	}
}

func (x *Queue) Push(job Job) {
	x.ch <- job
}

// Wait goroutine 으로 메인문에서 실행해준다.
func (x *Queue) Wait() {
	for job := range x.ch {
		var ctx = x.cg(context.TODO())
		var err = job(ctx)
		if err != nil {
			fnLogger.GetLogger(ctx).CtxError(
				ctx,
				map[string]any{
					"error": err.Error(),
				},
			)
		}
	}
}
