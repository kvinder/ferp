package view

import "net/http"

// Index render view
func Index(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("index.html"), w, data)
}

// Login render view
func Login(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("login.html"), w, data)
}

// FamsDashboard render view
func FamsDashboard(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("fams/dashboard.html"), w, data)
}

// FamsRequest render view
func FamsRequest(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("fams/request.html"), w, data)
}

// AdminCreateUser render view
func AdminCreateUser(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/create.html"), w, data)
}

// AdminList render view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(parseTemplate("admin/list.html"), w, data)
}
