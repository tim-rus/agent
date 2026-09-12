package dialog

import (
	"context"
	"core/internal/pg/chat"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
)

func (d *Dialog) Ask(ctx context.Context, req *connect.Request[chat.AskRequest]) (*connect.Response[chat.AskResponse], error) {
	res, err := d.Prompt(ctx, req.Msg.Prompt)
	if err != nil {
		slog.Error("failed to ask", "err", err)
		return nil, fmt.Errorf("request error")
	}
	return &connect.Response[chat.AskResponse]{
		Msg: &chat.AskResponse{
			Completion: res.Content,
		},
	}, nil
}
