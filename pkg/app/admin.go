package app

import (
	"ferp/pkg/model"
	"ferp/pkg/view"
	"net/http"
	"strconv"
	"time"
)

func adminCreate(w http.ResponseWriter, r *http.Request) {
	userOnLogin, err := model.UserOnLogin(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !checkRole("Admin", userOnLogin.Roles) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		user := model.User{
			EmployeeID: r.FormValue("field-employeeId"),
			Username:   r.FormValue("field-username"),
			Password:   r.FormValue("field-password"),
			Name:       r.FormValue("field-name"),
			Sex:        r.FormValue("field-sex"),
			Department: r.FormValue("field-department"),
			Email:      r.FormValue("field-email"),
			Telephone:  r.FormValue("field-telephone"),
			Roles:      r.Form["field-checkbox"],
			CreateDate: now,
			UpdateDate: now,
		}
		model.CreateUser(&user)
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	data := map[string]interface{}{
		"admin":        "active open",
		"admin_create": "active",
		"nameLogin":    userOnLogin.Name,
		"list_roles":   model.ListRoles(),
	}
	data = setAut(data, userOnLogin.Roles)
	view.AdminCreateUser(w, data)
}

func adminUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")
		id, _ := strconv.Atoi(r.FormValue("field-ID"))
		user := model.User{
			ID:         id,
			EmployeeID: r.FormValue("field-employeeId"),
			Username:   r.FormValue("field-username"),
			Password:   r.FormValue("field-password"),
			Name:       r.FormValue("field-name"),
			Sex:        r.FormValue("field-sex"),
			Department: r.FormValue("field-department"),
			Email:      r.FormValue("field-email"),
			Telephone:  r.FormValue("field-telephone"),
			Roles:      r.Form["field-checkbox"],
			CreateDate: now,
			UpdateDate: now,
		}
		model.UpdateUser(&user)
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
}

func adminList(w http.ResponseWriter, r *http.Request) {
	userOnLogin, err := model.UserOnLogin(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !checkRole("Admin", userOnLogin.Roles) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data := map[string]interface{}{
		"admin":      "active open",
		"admin_list": "active",
		"user_list":  model.ListUsers(),
		"nameLogin":  userOnLogin.Name,
	}
	data = setAut(data, userOnLogin.Roles)
	view.AdminList(w, data)
}

func adminUser(w http.ResponseWriter, r *http.Request) {
	userOnLogin, err := model.UserOnLogin(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := map[string]interface{}{
		"admin":      "active open",
		"list_roles": model.ListRoles(),
		"nameLogin":  userOnLogin.Name,
	}
	if !checkRole("Admin", userOnLogin.Roles) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data = setAut(data, userOnLogin.Roles)
	viewID := r.URL.Query().Get("view")
	updateID := r.URL.Query().Get("update")

	if len(viewID) != 0 && len(updateID) == 0 {
		data["view"] = model.GetUser(viewID)
		view.AdminUser(w, data)
		return
	}
	if len(viewID) == 0 && len(updateID) != 0 {
		data["update"] = model.GetUser(updateID)
		view.AdminUpdateUser(w, data)
		return
	}
	http.NotFound(w, r)
}
