package fnGoroutine

import (
	"context"
)

type ContextGenerator func(ctx context.Context) context.Context
type Job func(ctx context.Context) (err error)

const (
	ErrAlreadyInProgress = "already_in_progress"
	ErrNotInProgress     = "not_in_progress"
)
