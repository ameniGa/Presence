package context

import (
	"context"
	"testing"
)

func TestAddTimeoutToCtx(t *testing.T) {
	t.Run("timeout < 1", func(t *testing.T) {
		ctx, _ := AddTimeoutToCtx(context.Background(), 1)
		if ctx.Err() != nil {
			t.Errorf("Expected Successful, recieved: %v\n", ctx.Err().Error())
		}
	})

	t.Run("timeout > 1", func(t *testing.T) {
		ctx, _ := AddTimeoutToCtx(context.Background(), 2)
		if ctx.Err() != nil {
			t.Errorf("Expected Successful, recieved: %v\n", ctx.Err().Error())
		}
	})
}