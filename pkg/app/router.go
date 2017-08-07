package app

import (
	"net/http"
)

// Router handlers to mux
func Router(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(http.Dir("static"))))
	mux.Handle("/login", midelware(http.HandlerFunc(login)))

	mux.HandleFunc("/fa/dashboard", famsDashboard)

	//Admin
	needLoginAdminMux := http.NewServeMux()
	needLoginAdminMux.HandleFunc("/create", adminCreate)
	needLoginAdminMux.HandleFunc("/list", adminList)
	needLoginAdminMux.HandleFunc("/user", adminUser)
	mux.Handle("/admin/", http.StripPrefix("/admin", midelware(needLoginAdminMux)))

	//Need Login
	needLoginFaMux := http.NewServeMux()
	needLoginFaMux.HandleFunc("/request", famsRequest)
	mux.Handle("/fams/", http.StripPrefix("/fams", midelware(needLoginFaMux)))
}

func midelware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
