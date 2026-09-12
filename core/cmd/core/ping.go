package main

import (
	"context"
	"core/internal/pg/ping"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
)

type PingSvc struct {
}

func NewPing() *PingSvc {
	return &PingSvc{}
}

//

func (p *PingSvc) Ping(ctx context.Context, req *connect.Request[ping.PingRequest]) (*connect.Response[ping.PingResponse], error) {
	slog.Debug("PING", "msg", req.Msg.Msg)
	res := connect.NewResponse(&ping.PingResponse{
		Msg: fmt.Sprintf("msg received: %s", req.Msg),
	})
	return res, nil
}
