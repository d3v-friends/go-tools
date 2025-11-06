package fnGoroutine

import (
	"context"
	"sync"
	"time"

	"github.com/d3v-friends/go-tools/fnError"
	"github.com/d3v-friends/go-tools/fnLogger"
)

type SingleWorker struct {
	ch chan Job
	mu *sync.Mutex

	// context
	contextGenerator ContextGenerator
	fnCancel         context.CancelFunc

	// status
	inProgress    bool
	lastUpdatedAt time.Time

	// logger
	logGroup fnLogger.LogGroup
}

func NewSingleWorker(
	cg ContextGenerator,
	logGroup fnLogger.LogGroup,
) *SingleWorker {
	return &SingleWorker{
		ch:               make(chan Job),
		mu:               &sync.Mutex{},
		contextGenerator: cg,
		fnCancel:         nil,
		inProgress:       false,
		lastUpdatedAt:    time.Now(),
		logGroup:         logGroup,
	}
}

func (x *SingleWorker) Run(job Job) (err error) {
	x.mu.Lock()
	defer x.mu.Unlock()

	if x.inProgress {
		err = fnError.New(ErrAlreadyInProgress)
		return
	}

	x.ch <- job
	return
}

func (x *SingleWorker) Cancel() (err error) {
	x.mu.Lock()
	defer x.mu.Unlock()

	if !x.inProgress {
		err = fnError.New(ErrNotInProgress)
		return
	}

	x.fnCancel()
	x.inProgress = false
	x.fnCancel = nil
	x.lastUpdatedAt = time.Now()
	return
}

func (x *SingleWorker) Status() (inProgress bool, lastUpdatedAt time.Time) {
	x.mu.Lock()
	defer x.mu.Unlock()
	return x.inProgress, x.lastUpdatedAt
}

func (x *SingleWorker) Wait() {
	for {
		var job = <-x.ch
		var startAt = time.Now()
		var ctx = x.contextGenerator(context.TODO())

		x.mu.Lock()
		ctx, x.fnCancel = context.WithCancel(ctx)
		x.inProgress = true
		x.lastUpdatedAt = time.Now()
		x.mu.Unlock()

		var logger = fnLogger.GetLogger(ctx)

		var err = job(ctx)
		if err != nil {
			logger.CtxError(
				ctx,
				map[string]any{
					"error":    err.Error(),
					"duration": time.Since(startAt).Milliseconds(),
				},
				x.logGroup,
			)
		} else {
			logger.CtxInfo(
				ctx,
				map[string]any{
					"work":     "success",
					"duration": time.Since(startAt).Milliseconds(),
				},
				x.logGroup,
			)
		}

		x.mu.Lock()
		x.inProgress = false
		x.fnCancel = nil
		x.lastUpdatedAt = time.Now()
		x.mu.Unlock()
	}
}
