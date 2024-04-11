package main

import (
	"context"
	"syscall"

	"github.com/Garik-/ws-epoll/internal/app"
	"github.com/Garik-/ws-epoll/internal/closer"
	"github.com/Garik-/ws-epoll/internal/zlog"
	"golang.org/x/sync/errgroup"
)

func main() {
	defer zlog.Sync()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errGroup, errCtx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		return closer.CloseOnSignal(errCtx, syscall.SIGINT, syscall.SIGTERM)
	})

	errGroup.Go(func() error {
		srv := app.New(&app.Config{
			Addr: ":3133",
		})
		closer.Add(srv)

		return srv.Run(errCtx)
	})

	err := errGroup.Wait()
	if err != nil {
		zlog.Error("errGroup.Wait", zlog.Err(err))
	}
}
