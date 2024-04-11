package app

import (
	"net/http"

	"github.com/gobwas/ws"
)

func wsUpgrade(w http.ResponseWriter, r *http.Request) {
	_, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}
