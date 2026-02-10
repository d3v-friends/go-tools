package fnCron

import (
	"context"
	"sync"
	"time"

	"github.com/d3v-friends/go-tools/fnLogger"
)

type Worker struct {
	ch               *time.Ticker
	contextGenerator ContextGenerator
	jobs             []Job
}

func NewWorker(cg ContextGenerator, jobs ...Job) *Worker {
	return &Worker{
		ch:               time.NewTicker(time.Second),
		contextGenerator: cg,
		jobs:             jobs,
	}
}

func (x *Worker) Wait() {
	for now := range x.ch.C {
		var ctx = x.contextGenerator(context.TODO())
		var wg = &sync.WaitGroup{}
		wg.Add(len(x.jobs))

		for _, job := range x.jobs {
			if !job.IsRun(now) {
				continue
			}
			go do(ctx, job, wg)
		}

		wg.Wait()
	}

}

func do(ctx context.Context, job Job, wg *sync.WaitGroup) {
	defer wg.Done()

	var startAt = time.Now()
	var logGroup = fnLogger.NewLogGroup(job.GetName(), fnLogger.ColorKeyYellow)
	var logger = fnLogger.GetLogger(ctx)
	var err = job.Do(ctx)

	if err != nil {
		logger.CtxError(
			ctx,
			map[string]any{
				"error":    err.Error(),
				"duration": time.Since(startAt).Milliseconds(),
			},
			logGroup,
		)
		return
	}

	logger.CtxInfo(
		ctx,
		map[string]any{
			"duration": time.Since(startAt).Milliseconds(),
		},
		logGroup,
	)
}
