package app

import (
	"ferp/pkg/view"
	"net/http"
)

func famsDashboard(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"fams":           "active open",
		"fams_dashboard": "active",
	}
	view.FamsDashboard(w, data)
}

func famsRequest(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"fams":         "active open",
		"fams_request": "active",
	}
	view.FamsRequest(w, data)
}
