package app

import (
	"ferp/pkg/model"
	"ferp/pkg/view"
	"net/http"
	"time"
)

func adminCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		user := model.User{
			EmployeeID: r.FormValue("field-employeeId"),
			Username:   r.FormValue("field-username"),
			Password:   r.FormValue("field-password"),
			Name:       r.FormValue("field-name"),
			Sex:        r.FormValue("field-sex"),
			Department: r.FormValue("field-department"),
			Email:      r.FormValue("field-email"),
			Telephone:  r.FormValue("field-telephone"),
			Role:       r.Form["field-checkbox"],
			CreateDate: time.Now(),
			UpdateDate: time.Now(),
		}
		model.CreateUser(&user)
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"admin":        "active open",
		"admin_create": "active",
	}
	view.AdminCreateUser(w, data)
}

func adminList(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"admin":      "active open",
		"admin_list": "active",
		"user_list":  model.ListUsers(),
	}
	view.AdminList(w, data)
}

func adminUser(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"admin":      "active",
		"admin_list": "active",
		"user_list":  model.ListUsers(),
	}
	view.AdminList(w, data)
}
