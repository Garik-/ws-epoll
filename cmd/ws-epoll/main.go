package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/Garik-/ws-epoll/internal/app"
	"github.com/Garik-/ws-epoll/internal/closer"
	"github.com/Garik-/ws-epoll/internal/zlog"
	"golang.org/x/sync/errgroup"
)

const (
	appName     = "ws-epoll"
	defaultAddr = ":3133"
)

var (
	fAddr string
	fHelp bool
)

func setupFlags() {
	flag.StringVar(&fAddr, "addr", defaultAddr, "server address")
	flag.BoolVar(&fHelp, "help", false, "show this help")
	flag.Usage = usage
}

func usage() {
	var sb strings.Builder

	fmt.Fprintf(&sb, "%s - test 1 million connections, usage:\n\n", appName)
	fmt.Fprintf(&sb, "%s [flags]\n\n", filepath.Base(os.Args[0]))
	fmt.Fprint(&sb, "possible flags with default values:\n\n")

	_, _ = os.Stderr.WriteString(sb.String())

	flag.PrintDefaults()
}

func main() {
	defer zlog.Sync()

	setupFlags()
	flag.Parse()

	if fHelp {
		usage()

		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errGroup, errCtx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		return closer.CloseOnSignal(errCtx, syscall.SIGINT, syscall.SIGTERM)
	})

	errGroup.Go(func() error {
		srv := app.New(&app.Config{
			Addr: fAddr,
		})
		closer.Add(srv)

		return srv.Run(errCtx)
	})

	err := errGroup.Wait()
	if err != nil {
		zlog.Error("errGroup.Wait", zlog.Err(err))
	}
}
