package main

import (
	"context"
	pingGen "core/internal/gen"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
)

type Ping struct{}

//

func (api *Ping) Ping(_ context.Context, req *connect.Request[pingGen.PingRequest]) (*connect.Response[pingGen.PingResponse], error) {
	slog.Info("msg received", "msg", req.Msg)
	res := connect.NewResponse(&pingGen.PingResponse{
		Msg: fmt.Sprintf("msg received: %s", req.Msg),
	})
	return res, nil
}
