package protocol

import "context"

func RunHandshakeDeadline2(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	svc := HandshakeDeadline2Service{}
	return svc.Execute(ctx, func(callCtx context.Context) error {
		if err := callCtx.Err(); err != nil {
			return err
		}
		return nil
	})
}
