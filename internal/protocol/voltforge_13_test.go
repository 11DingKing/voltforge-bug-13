package protocol

import (
	"context"
	"errors"
	"testing"
)

func TestVoltForge13(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := RunHandshakeDeadline2(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
}
