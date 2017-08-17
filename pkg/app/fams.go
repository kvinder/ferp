package app

import (
	"ferp/pkg/model"
	"ferp/pkg/view"
	"net/http"
)

func famsDashboard(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"fams":           "active open",
		"fams_dashboard": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	view.FamsDashboard(w, data)
}

func famsRequest(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"fams":         "active open",
		"fams_request": "active",
	}
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	if !checkRoles([]string{"Admin", "Sale_Co"}, userOnLogin.Roles) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	view.FamsRequest(w, data)
}
