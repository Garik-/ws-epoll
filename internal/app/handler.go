package app

import (
	"net/http"

	"github.com/Garik-/ws-epoll/internal/zlog"
)

func wsUpdate(_ http.ResponseWriter, _ *http.Request) {
	zlog.Debug("trlolo")
}
