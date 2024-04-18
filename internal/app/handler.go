package app

import (
	"net/http"
)

func wsUpgrade(w http.ResponseWriter, _ *http.Request) {
	// _, _, _, err := ws.UpgradeHTTP(r, w)
	// if err != nil {
	http.Error(w, http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}
