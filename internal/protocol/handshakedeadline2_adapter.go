package protocol

import "context"

func RunHandshakeDeadline2(ctx context.Context) error {
	svc := HandshakeDeadline2Service{}
	return svc.Execute(ctx, func(callCtx context.Context) error {
		return callCtx.Err()
	})
}
