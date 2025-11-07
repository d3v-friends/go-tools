package fnCron

import (
	"context"
	"time"

	"github.com/d3v-friends/go-tools/fnLogger"
)

type SingleWorker struct {
	ch               *time.Ticker
	contextGenerator ContextGenerator
	job              Job
	logGroup         fnLogger.LogGroup
}

type ContextGenerator func(ctx context.Context) context.Context

type Job interface {
	Do(ctx context.Context) (err error)
	IsRun(now time.Time) bool
	GetName() string
}

func NewSingleWorker(cg ContextGenerator, job Job) *SingleWorker {
	return &SingleWorker{
		ch:               time.NewTicker(time.Second),
		contextGenerator: cg,
		job:              job,
		logGroup:         fnLogger.NewLogGroup(job.GetName(), fnLogger.ColorKeyYellow),
	}
}

// Wait
// main 문에서 goroutine 으로 대기 시킨다.
func (x *SingleWorker) Wait() {
	for now := range x.ch.C {
		if !x.job.IsRun(now) {
			continue
		}
		x.Run()
	}
}

func (x *SingleWorker) Run() {
	var startAt = time.Now()
	var ctx = x.contextGenerator(context.TODO())
	var err = x.job.Do(ctx)

	if err != nil {
		fnLogger.GetLogger(ctx).CtxError(
			ctx,
			map[string]any{
				"job":      x.job.GetName(),
				"duration": time.Since(startAt).Milliseconds(),
				"err":      err.Error(),
			},
			x.logGroup,
		)
		return
	}

	fnLogger.GetLogger(ctx).CtxInfo(
		ctx,
		map[string]any{
			"job":      x.job.GetName(),
			"duration": time.Since(startAt).Milliseconds(),
		},
		x.logGroup,
	)
}
