package protocol

import (
	"context"
	"fmt"
)

type HandshakeDeadline2Service struct{}

func (s HandshakeDeadline2Service) Execute(ctx context.Context, action func(context.Context) error) error {
	if ctx == nil {
		return fmt.Errorf("handshakedeadline2: nil context")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("handshakedeadline2 before action: %w", err)
	}
	if err := action(ctx); err != nil {
		return fmt.Errorf("handshakedeadline2 action: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("handshakedeadline2 after action: %w", err)
	}
	return nil
}
