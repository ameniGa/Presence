package context

import (
	"context"
	"time"
)


// AddTimeoutToCtx add timeout to context.
func AddTimeoutToCtx(ctx context.Context, timeOutDuration int64) (context.Context, context.CancelFunc) {
	timeOut := time.Duration(timeOutDuration) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	return ctx, cancel
}