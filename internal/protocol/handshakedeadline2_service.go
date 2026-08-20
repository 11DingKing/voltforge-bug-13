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
	return action(context.Background())
}
