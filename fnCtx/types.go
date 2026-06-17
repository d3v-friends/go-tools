package fnCtx

import "context"

type (
	ContextGenerator func(ctx context.Context) context.Context
)
