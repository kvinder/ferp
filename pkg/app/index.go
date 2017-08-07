package app

import (
	"ferp/pkg/view"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"homePage": "active",
	}
	view.Index(w, data)
}
