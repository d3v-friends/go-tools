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

	var err = job(ctx)
	if err != nil {
		x.appendError(err)
	}

	x.incInProgress(-1)
	signal <- time.Now()
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

/* ---------------------------------------------------------------------- */

type Result[T any] struct {
	Result T
	Error  error
}

type FnThreadWithResult[T any] func(ctx context.Context) (res T, err error)

type ThreadWithResult[T any] struct {
	size       int
	inProgress int
	status     []*Status
	mu         *sync.Mutex
	result     []*Result[T]
}

func NewThreadWithResult[T any](size int) *ThreadWithResult[T] {
	return &ThreadWithResult[T]{
		size:       size,
		inProgress: 0,
		status:     make([]*Status, 0),
		mu:         &sync.Mutex{},
		result:     make([]*Result[T], 0),
	}
}

func (x *ThreadWithResult[T]) Do(
	ctx context.Context,
	jobs []FnThreadWithResult[T],
) (ls []*Result[T]) {
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

	return x.result
}

func (x *ThreadWithResult[T]) work(
	ctx context.Context,
	job FnThreadWithResult[T],
	signal chan<- time.Time,
) {
	x.incInProgress(1)

	var res, err = job(ctx)

	x.appendResult(&Result[T]{
		Result: res,
		Error:  err,
	})

	x.incInProgress(-1)
	signal <- time.Now()
}

func (x *ThreadWithResult[T]) incInProgress(i int) {
	x.mu.Lock()
	x.inProgress += i
	x.status = append(x.status, &Status{
		Now:        time.Now(),
		InProgress: x.inProgress,
	})
	x.mu.Unlock()
}

func (x *ThreadWithResult[T]) appendResult(res *Result[T]) {
	x.mu.Lock()
	x.result = append(x.result, res)
	x.mu.Unlock()
}
