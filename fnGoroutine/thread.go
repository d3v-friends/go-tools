package fnGoroutine

import (
	"context"
	"sync"
	"time"
)

type Status struct {
	Now        time.Time
	InProgress int
}

type Thread struct {
	size       int
	inProgress int
	status     []*Status
	mu         *sync.Mutex
	errors     []error
}

func NewThread(size int) *Thread {
	return &Thread{
		size:       size,
		inProgress: 0,
		mu:         &sync.Mutex{},
		errors:     make([]error, 0),
		status:     make([]*Status, 0),
	}
}

func (x *Thread) Do(
	ctx context.Context,
	jobs []Job,
) (errs []error) {
	if len(jobs) <= x.size {
		x.size = len(jobs)
	}

	var signal = make(chan time.Time)

	for i := 0; i < x.size; i++ {
		go x.work(ctx, jobs[i], signal)
	}

	for i := x.size; i < len(jobs); i++ {
		<-signal
		go x.work(ctx, jobs[i], signal)
	}

	for i := 0; i < x.size; i++ {
		<-signal
	}

	if len(x.errors) == 0 {
		return nil
	}

	return x.errors
}

func (x *Thread) work(
	ctx context.Context,
	job Job,
	signal chan<- time.Time,
) {
	x.incInProgress(1)
	defer func() {
		x.incInProgress(-1)
		signal <- time.Now()
	}()

	var err = job(ctx)
	if err != nil {
		x.appendError(err)
	}
}

func (x *Thread) incInProgress(i int) {
	x.mu.Lock()
	x.inProgress += i
	x.status = append(x.status, &Status{
		Now:        time.Now(),
		InProgress: x.inProgress,
	})
	x.mu.Unlock()
}

func (x *Thread) appendError(err error) {
	x.mu.Lock()
	x.errors = append(x.errors, err)
	x.mu.Unlock()
}
