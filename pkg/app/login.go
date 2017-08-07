package app

import (
	"ferp/pkg/view"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	view.Login(w, nil)
}
