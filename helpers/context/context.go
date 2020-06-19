package context

import (
	"context"
	"errors"
	"time"
)


// AddTimeoutToCtx add timeout to context.
func AddTimeoutToCtx(ctx context.Context, timeOutDuration int64) (context.Context, context.CancelFunc) {
	timeOut := time.Duration(timeOutDuration) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	return ctx, cancel
}


func IsValidCtx(ctx context.Context, maxTimeout time.Duration) (bool, error) {
	// check timeout and sets default timeout if not specified
	if _, ok := ctx.Deadline(); !ok {
		return false, errors.New("missing context deadline")
	}
	if deadline, _ := ctx.Deadline(); deadline.Sub(time.Now()) > maxTimeout*time.Second {
		return false, errors.New("context deadline exceeded")
	}
	return true, nil
}