package app

import (
	"ferp/pkg/model"
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
	userOnLogin, err := model.UserOnLogin(r)
	if err == nil {
		data["nameLogin"] = userOnLogin.Name
		data = setAut(data, userOnLogin.Roles)
	}
	view.Index(w, data)
}
